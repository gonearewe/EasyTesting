package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gonearewe/EasyTesting/handlers"
)

func SetupRoute(UnauthRouter gin.IRouter, teacherAuthRouter *gin.RouterGroup, adminAuthRouter *gin.RouterGroup) {
	UnauthRouter.GET("/ping", handlers.PingHandler)
	adminAuthRouter.POST("/teachers", handlers.TeachersRegisterHandler)
	adminAuthRouter.GET("/teachers", handlers.GetTeachersHandler)
	teacherAuthRouter.GET("/hello", handlers.HelloHandler)
}
