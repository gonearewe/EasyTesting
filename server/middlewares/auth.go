package middlewares

import (
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gonearewe/EasyTesting/dao"
	"github.com/gonearewe/EasyTesting/models"
	"github.com/gonearewe/EasyTesting/utils"
	"github.com/spf13/viper"
)

func setupAuth(r *gin.Engine) (teacherAuthRouter *gin.RouterGroup, adminAuthRouter *gin.RouterGroup) {
	teacherAuthMiddleware := generateAuthMiddleware(teacherAuthenticator, teacherPayLoadFunc, nil)
	adminAuthMiddleware := generateAuthMiddleware(teacherAuthenticator, teacherPayLoadFunc, adminAuthorizator)
	r.GET("/teacher_auth", teacherAuthMiddleware.LoginHandler)
	teacherAuthRouter = r.Group("/")
	teacherAuthRouter.Use(teacherAuthMiddleware.MiddlewareFunc())
	adminAuthRouter = r.Group("/")
	adminAuthRouter.Use(adminAuthMiddleware.MiddlewareFunc())
	return
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
		Timeout:  	   72*time.Hour,
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
	id := c.Query("teacher_id")
	teacher := dao.GetTeacherByTeacherId(id)
	err = utils.VerifyPassword(c.Query("password"), teacher.Salt, teacher.Password)
	utils.PanicWhen(err)
	return teacher, nil
}

func teacherPayLoadFunc(data interface{}) jwt.MapClaims {
	v := data.(*models.Teacher)
	return jwt.MapClaims{
		"id":v.ID,
		"teacher_id": v.TeacherID,
		"is_admin":   v.IsAdmin,
	}
}

func adminAuthorizator(data interface{}, c *gin.Context) bool {
	return jwt.ExtractClaims(c)["is_admin"].(bool)
}
