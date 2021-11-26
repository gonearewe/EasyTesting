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
}
