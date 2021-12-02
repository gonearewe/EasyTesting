package middlewares

import "github.com/gin-gonic/gin"

func SetupMiddleWares(r *gin.Engine) (teacherAuthRoute *gin.RouterGroup, adminAuthRoute *gin.RouterGroup) {
	r.Use(gin.Logger())
	r.Use(recovery)
	return setupAuth(r)
}
