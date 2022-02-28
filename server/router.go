package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gonearewe/EasyTesting/handlers"
	"github.com/gonearewe/EasyTesting/middlewares"
)

func SetupRoute(unauthRouter gin.IRouter, teacherAuthRouter *gin.RouterGroup,
	adminAuthRouter *gin.RouterGroup, studentAuthRouter *gin.RouterGroup) {
	unauthRouter.StaticFile("/","./index.html")
	unauthRouter.StaticFile("/favicon.ico","./favicon.ico")
	unauthRouter.Static("/static/","./static")
	unauthRouter.GET("/ping", handlers.PingHandler)
	teacherAuthRouter.GET("/hello", handlers.HelloHandler)
	studentAuthRouter.GET("/cache",handlers.GetCacheHandler)
	studentAuthRouter.PUT("/cache",handlers.PutCacheHandler)

	adminAuthRouter.GET("/teachers", handlers.GetTeachersHandler)
	adminAuthRouter.POST("/teachers", handlers.TeachersRegisterHandler)
	adminAuthRouter.PUT("/teachers", handlers.PutTeacherHandler)
	adminAuthRouter.DELETE("/teachers",middlewares.CheckGlobalExamStatus, handlers.DeleteTeachersHandler)

	teacherAuthRouter.PUT("/profile", handlers.PutTeacherProfileHandler)

	teacherAuthRouter.GET("/students", handlers.GetStudentsHandler)
	teacherAuthRouter.POST("/students", handlers.StudentsRegisterHandler)
	teacherAuthRouter.PUT("/students",middlewares.CheckGlobalExamStatus, handlers.PutStudentHandler)
	teacherAuthRouter.DELETE("/students",middlewares.CheckGlobalExamStatus, handlers.DeleteStudentsHandler)

	teacherAuthRouter.GET("/mcq", handlers.GetMcqHandler)
	teacherAuthRouter.POST("/mcq",middlewares.CheckGlobalExamStatus, handlers.PostMcqHandler)
	teacherAuthRouter.PUT("/mcq",middlewares.CheckGlobalExamStatus, handlers.PutMcqHandler)
	teacherAuthRouter.DELETE("/mcq",middlewares.CheckGlobalExamStatus, handlers.DeleteMcqHandler)

	teacherAuthRouter.GET("/maq", handlers.GetMaqHandler)
	teacherAuthRouter.POST("/maq",middlewares.CheckGlobalExamStatus, handlers.PostMaqHandler)
	teacherAuthRouter.PUT("/maq",middlewares.CheckGlobalExamStatus, handlers.PutMaqHandler)
	teacherAuthRouter.DELETE("/maq",middlewares.CheckGlobalExamStatus, handlers.DeleteMaqHandler)

	teacherAuthRouter.GET("/bfq", handlers.GetBfqHandler)
	teacherAuthRouter.POST("/bfq",middlewares.CheckGlobalExamStatus, handlers.PostBfqHandler)
	teacherAuthRouter.PUT("/bfq",middlewares.CheckGlobalExamStatus, handlers.PutBfqHandler)
	teacherAuthRouter.DELETE("/bfq",middlewares.CheckGlobalExamStatus, handlers.DeleteBfqHandler)

	teacherAuthRouter.GET("/tfq", handlers.GetTfqHandler)
	teacherAuthRouter.POST("/tfq",middlewares.CheckGlobalExamStatus, handlers.PostTfqHandler)
	teacherAuthRouter.PUT("/tfq",middlewares.CheckGlobalExamStatus, handlers.PutTfqHandler)
	teacherAuthRouter.DELETE("/tfq",middlewares.CheckGlobalExamStatus, handlers.DeleteTfqHandler)

	teacherAuthRouter.GET("/crq", handlers.GetCrqHandler)
	teacherAuthRouter.POST("/crq",middlewares.CheckGlobalExamStatus, handlers.PostCrqHandler)
	teacherAuthRouter.PUT("/crq",middlewares.CheckGlobalExamStatus, handlers.PutCrqHandler)
	teacherAuthRouter.DELETE("/crq",middlewares.CheckGlobalExamStatus, handlers.DeleteCrqHandler)

	teacherAuthRouter.GET("/cq", handlers.GetCqHandler)
	teacherAuthRouter.POST("/cq",middlewares.CheckGlobalExamStatus, handlers.PostCqHandler)
	teacherAuthRouter.PUT("/cq",middlewares.CheckGlobalExamStatus, handlers.PutCqHandler)
	teacherAuthRouter.DELETE("/cq",middlewares.CheckGlobalExamStatus, handlers.DeleteCqHandler)

	teacherAuthRouter.GET("/exams", handlers.GetExamHandler)
	teacherAuthRouter.GET("/exams/ended", handlers.GetEndedExamHandler)
	teacherAuthRouter.POST("/exams", handlers.PostExamHandler)
	teacherAuthRouter.PUT("/exams", handlers.PutExamHandler)
	teacherAuthRouter.DELETE("/exams", handlers.DeleteExamHandler)

	teacherAuthRouter.GET("/exams/examinees", handlers.GetExamineeHandler)
	studentAuthRouter.GET("/exams/my_questions", handlers.GetMyQuestionsHandler)
	studentAuthRouter.PUT("/exams/my_answers", handlers.PutMyAnswersHandler)
}
