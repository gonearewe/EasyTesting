package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gonearewe/EasyTesting/handlers"
)

func SetupRoute(r gin.IRouter) {
	r.POST("/ping", handlers.PingHandler)
	r.POST("/teachers", handlers.TeachersRegisterHandler)
}
