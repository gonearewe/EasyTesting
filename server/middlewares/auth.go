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

func setupAuth(r *gin.Engine) (teacherAuthRouter *gin.RouterGroup,
	adminAuthRouter *gin.RouterGroup, studentAuthRouter *gin.RouterGroup) {
	teacherAuthMiddleware := generateAuthMiddleware(teacherAuthenticator, teacherPayLoadFunc, nil)
	adminAuthMiddleware := generateAuthMiddleware(teacherAuthenticator, teacherPayLoadFunc, adminAuthorizator)
	studentAuthMiddleware := generateAuthMiddleware(studentAuthenticator, studentPayLoadFunc, nil)
	r.GET("/teacher_auth", teacherAuthMiddleware.LoginHandler)
	teacherAuthRouter = r.Group("/")
	teacherAuthRouter.Use(teacherAuthMiddleware.MiddlewareFunc())
	adminAuthRouter = r.Group("/")
	adminAuthRouter.Use(adminAuthMiddleware.MiddlewareFunc())
	r.GET("/student_auth", studentAuthMiddleware.LoginHandler)
	studentAuthRouter = r.Group("/")
	studentAuthRouter.Use(studentAuthMiddleware.MiddlewareFunc())
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
		Timeout:       72 * time.Hour,
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
		"id":         v.ID,
		"teacher_id": v.TeacherID,
		"name":       v.Name,
		"is_admin":   v.IsAdmin,
	}
}

func adminAuthorizator(data interface{}, c *gin.Context) bool {
	return jwt.ExtractClaims(c)["is_admin"].(bool)
}

func studentAuthenticator(c *gin.Context) (user interface{}, err error) {
	defer func() {
		if recover() != nil {
			user, err = nil, jwt.ErrFailedAuthentication
		}
	}()

	studentId := c.Query("student_id")
	examId := utils.Int(c.Query("exam_id"))
	student := dao.GetStudentBy(studentId)

	var session *models.ExamSession
	if err, session = dao.GetExamSessionBy(studentId, examId); err != nil {
		// try entering exam first
		dao.EnterExam(studentId, examId)
		if err, session = dao.GetExamSessionBy(studentId, examId); err != nil {
			// still fail
			utils.PanicWhen(err)
		}
		// first time get exam_session, we enter exam first and succeed
	}
	// else we've already entered the exam, and can get exam_session directly

	return jwt.MapClaims{
		"student_id":      studentId,
		"name":            student.Name,
		"class_id":        student.ClassID,
		"exam_session_id": session.ID,
	}, nil
}

func studentPayLoadFunc(data interface{}) jwt.MapClaims {
	return data.(jwt.MapClaims)
}
