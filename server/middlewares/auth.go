package middlewares

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gonearewe/EasyTesting/dao"
	"github.com/gonearewe/EasyTesting/handlers"
	"github.com/gonearewe/EasyTesting/models"
	"github.com/gonearewe/EasyTesting/utils"
	"github.com/spf13/viper"
)

func setupAuth(r *gin.Engine) {
	teacherAuthMiddleware := generateAuthMiddleware(teacherAuthenticator, teacherPayLoadFunc, nil)
	adminAuthMiddleware := generateAuthMiddleware(teacherAuthenticator, teacherPayLoadFunc, adminAuthorizator)
	r.GET("/teacher_auth", teacherAuthMiddleware.LoginHandler)
	authRequired := r.Group("/")
	authRequired.Use(teacherAuthMiddleware.MiddlewareFunc())
	adminRequired := authRequired.Group("admin")
	adminRequired.Use(adminAuthMiddleware.MiddlewareFunc())
	adminRequired.GET("/teachers", handlers.GetTeachersHandler)
	authRequired.POST("/hello", handlers.HelloHandler)
}

func generateAuthMiddleware(
	authenticator func(c *gin.Context) (interface{}, error),
	payloadFunc func(data interface{}) jwt.MapClaims,
	authorizator func(data interface{}, c *gin.Context) bool,
) (ret *jwt.GinJWTMiddleware) {
	var err error
	ret, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "easy testing",
		Key:           []byte(viper.GetString("jwt_secret_key")),
		Authenticator: authenticator,
		PayloadFunc:   payloadFunc,
		Authorizator:  authorizator,
	})
	utils.PanicWhen(err)
	return
}

func teacherAuthenticator(c *gin.Context) (user interface{}, err error) {
	defer func() {
		if recover() != nil {
			user, err = nil, jwt.ErrFailedAuthentication
		}
	}()
	json := utils.MustParseJson(c)
	id := json["teacher_id"].(string)
	teacher := dao.GetTeacherByTeacherId(id)
	err = utils.VerifyPassword(json["password"].(string), teacher.Salt, teacher.Password)
	utils.PanicWhen(err)
	return teacher, nil
}

func teacherPayLoadFunc(data interface{}) jwt.MapClaims {
	v := data.(*models.Teacher)
	return jwt.MapClaims{
		"teacher_id": v.TeacherID,
		"is_admin":   v.IsAdmin,
	}
}

func adminAuthorizator(data interface{}, c *gin.Context) bool {
	return jwt.ExtractClaims(c)["is_admin"].(bool)
}
