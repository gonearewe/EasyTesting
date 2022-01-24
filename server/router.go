package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gonearewe/EasyTesting/handlers"
)

func SetupRoute(unauthRouter gin.IRouter, teacherAuthRouter *gin.RouterGroup,
	adminAuthRouter *gin.RouterGroup, studentAuthRouter *gin.RouterGroup) {
	unauthRouter.StaticFile("/","../student_client/vue/dist/index.html")
	unauthRouter.Static("/static/","../student_client/vue/dist/static")
	unauthRouter.GET("/ping", handlers.PingHandler)
	teacherAuthRouter.GET("/hello", handlers.HelloHandler)
	studentAuthRouter.GET("/cache",handlers.GetCacheHandler)
	studentAuthRouter.PUT("/cache",handlers.PutCacheHandler)

	adminAuthRouter.GET("/teachers", handlers.GetTeachersHandler)
	adminAuthRouter.POST("/teachers", handlers.TeachersRegisterHandler)
	adminAuthRouter.PUT("/teachers", handlers.PutTeacherHandler)
	adminAuthRouter.DELETE("/teachers", handlers.DeleteTeachersHandler)

	teacherAuthRouter.PUT("/profile", handlers.PutTeacherProfileHandler)

	teacherAuthRouter.GET("/students", handlers.GetStudentsHandler)
	teacherAuthRouter.POST("/students", handlers.StudentsRegisterHandler)
	teacherAuthRouter.PUT("/students", handlers.PutStudentHandler)
	teacherAuthRouter.DELETE("/students", handlers.DeleteStudentsHandler)

	teacherAuthRouter.GET("/mcq", handlers.GetMcqHandler)
	teacherAuthRouter.POST("/mcq", handlers.PostMcqHandler)
	teacherAuthRouter.PUT("/mcq", handlers.PutMcqHandler)
	teacherAuthRouter.DELETE("/mcq", handlers.DeleteMcqHandler)

	teacherAuthRouter.GET("/maq", handlers.GetMaqHandler)
	teacherAuthRouter.POST("/maq", handlers.PostMaqHandler)
	teacherAuthRouter.PUT("/maq", handlers.PutMaqHandler)
	teacherAuthRouter.DELETE("/maq", handlers.DeleteMaqHandler)

	teacherAuthRouter.GET("/bfq", handlers.GetBfqHandler)
	teacherAuthRouter.POST("/bfq", handlers.PostBfqHandler)
	teacherAuthRouter.PUT("/bfq", handlers.PutBfqHandler)
	teacherAuthRouter.DELETE("/bfq", handlers.DeleteBfqHandler)

	teacherAuthRouter.GET("/tfq", handlers.GetTfqHandler)
	teacherAuthRouter.POST("/tfq", handlers.PostTfqHandler)
	teacherAuthRouter.PUT("/tfq", handlers.PutTfqHandler)
	teacherAuthRouter.DELETE("/tfq", handlers.DeleteTfqHandler)

	teacherAuthRouter.GET("/crq", handlers.GetCrqHandler)
	teacherAuthRouter.POST("/crq", handlers.PostCrqHandler)
	teacherAuthRouter.PUT("/crq", handlers.PutCrqHandler)
	teacherAuthRouter.DELETE("/crq", handlers.DeleteCrqHandler)

	teacherAuthRouter.GET("/cq", handlers.GetCqHandler)
	teacherAuthRouter.POST("/cq", handlers.PostCqHandler)
	teacherAuthRouter.PUT("/cq", handlers.PutCqHandler)
	teacherAuthRouter.DELETE("/cq", handlers.DeleteCqHandler)

	teacherAuthRouter.GET("/exams", handlers.GetExamHandler)
	teacherAuthRouter.GET("/exams/ended", handlers.GetEndedExamHandler)
	teacherAuthRouter.POST("/exams", handlers.PostExamHandler)
	teacherAuthRouter.PUT("/exams", handlers.PutExamHandler)
	teacherAuthRouter.DELETE("/exams", handlers.DeleteExamHandler)

	teacherAuthRouter.GET("/exams/examinees", handlers.GetExamineeHandler)
	studentAuthRouter.GET("/exams/my_questions", handlers.GetMyQuestionsHandler)
	studentAuthRouter.PUT("/exams/my_answers", handlers.PutMyAnswersHandler)
}
