package middlewares

import (
	"github.com/gin-gonic/gin"
)

func SetupMiddleWares(r *gin.Engine) (teacherAuthRoute *gin.RouterGroup,
	adminAuthRoute *gin.RouterGroup, studentAuthRouter *gin.RouterGroup) {
	r.Use(logMiddleware)
	r.Use(gin.Logger())
	r.Use(recovery)
	r.Use(corsMiddleware())
	return setupAuth(r)
}
