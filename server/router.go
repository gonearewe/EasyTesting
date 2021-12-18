package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gonearewe/EasyTesting/handlers"
)

func SetupRoute(unauthRouter gin.IRouter, teacherAuthRouter *gin.RouterGroup, adminAuthRouter *gin.RouterGroup) {
    unauthRouter.GET("/ping", handlers.PingHandler)
    teacherAuthRouter.GET("/hello", handlers.HelloHandler)

    adminAuthRouter.GET("/teachers", handlers.GetTeachersHandler)
    adminAuthRouter.POST("/teachers", handlers.TeachersRegisterHandler)
    adminAuthRouter.PUT("/teachers",handlers.PutTeacherHandler)
    adminAuthRouter.DELETE("/teachers",handlers.DeleteTeachersHandler)

    teacherAuthRouter.GET("/students",handlers.GetStudentsHandler)
    teacherAuthRouter.POST("/students", handlers.StudentsRegisterHandler)
    teacherAuthRouter.PUT("/students",handlers.PutStudentHandler)
    teacherAuthRouter.DELETE("/students",handlers.DeleteStudentsHandler)

    teacherAuthRouter.GET("/mcq", handlers.GetMcqHandler)
    teacherAuthRouter.POST("/mcq", handlers.PostMcqHandler)
    teacherAuthRouter.PUT("/mcq", handlers.PutMcqHandler)
    teacherAuthRouter.DELETE("/mcq", handlers.DeleteMcqHandler)

    teacherAuthRouter.GET("/maq", handlers.GetMaqHandler)
    teacherAuthRouter.POST("/maq", handlers.PostMaqHandler)
    teacherAuthRouter.PUT("/maq", handlers.PutMaqHandler)
    teacherAuthRouter.DELETE("/maq", handlers.DeleteMaqHandler)

    teacherAuthRouter.GET("/exams", handlers.GetExamHandler)
    teacherAuthRouter.POST("/exams", handlers.PostExamHandler)
    teacherAuthRouter.PUT("/exams", handlers.PutExamHandler)
    teacherAuthRouter.DELETE("/exams", handlers.DeleteExamHandler)
}
