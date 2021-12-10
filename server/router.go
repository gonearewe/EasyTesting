package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gonearewe/EasyTesting/handlers"
)

func SetupRoute(unauthRouter gin.IRouter, teacherAuthRouter *gin.RouterGroup, adminAuthRouter *gin.RouterGroup) {
    unauthRouter.GET("/ping", handlers.PingHandler)
    teacherAuthRouter.GET("/hello", handlers.HelloHandler)

    adminAuthRouter.GET("/teachers/num", handlers.GetTeacherNumHandler)
    adminAuthRouter.GET("/teachers", handlers.GetTeachersHandler)
    adminAuthRouter.POST("/teachers", handlers.TeachersRegisterHandler)
    adminAuthRouter.PUT("/teachers",handlers.PutTeacherHandler)
    adminAuthRouter.DELETE("/teachers",handlers.DeleteTeacherHandler)

    teacherAuthRouter.GET("/students",handlers.GetStudentsHandler)
    teacherAuthRouter.GET("/students/num",handlers.GetStudentNumHandler)
    teacherAuthRouter.POST("/students", handlers.StudentsRegisterHandler)
    teacherAuthRouter.PUT("/students",handlers.PutStudentHandler)
    teacherAuthRouter.DELETE("/students",handlers.DeleteStudentHandler)

    teacherAuthRouter.GET("/mcq", handlers.GetMcqHandler)
    teacherAuthRouter.POST("/mcq", handlers.PostMcqHandler)
    teacherAuthRouter.PUT("/mcq", handlers.PutMcqHandler)
    teacherAuthRouter.DELETE("/mcq", handlers.DeleteMcqHandler)

    teacherAuthRouter.GET("/maq", handlers.GetMaqHandler)
    teacherAuthRouter.POST("/maq", handlers.PostMaqHandler)
    teacherAuthRouter.PUT("/maq", handlers.PutMaqHandler)
    teacherAuthRouter.DELETE("/maq", handlers.DeleteMaqHandler)

    teacherAuthRouter.GET("/exam", handlers.GetExamHandler)
    teacherAuthRouter.POST("/exam", handlers.PostExamHandler)
    teacherAuthRouter.PUT("/exam", handlers.PutExamHandler)
    teacherAuthRouter.DELETE("/exam", handlers.DeleteExamHandler)
}
