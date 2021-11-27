package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gonearewe/EasyTesting/handlers"
)

func SetupRoute(unauthRouter gin.IRouter, teacherAuthRouter *gin.RouterGroup, adminAuthRouter *gin.RouterGroup) {
    unauthRouter.GET("/ping", handlers.PingHandler)
    adminAuthRouter.POST("/teachers", handlers.TeachersRegisterHandler)
    adminAuthRouter.GET("/teachers", handlers.GetTeachersHandler)
    teacherAuthRouter.GET("/hello", handlers.HelloHandler)

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
