package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupMiddleWares(r *gin.Engine) (teacherAuthRoute *gin.RouterGroup,
	adminAuthRoute *gin.RouterGroup, studentAuthRouter *gin.RouterGroup) {
	r.Use(gin.Logger())
	r.Use(recovery)
	r.Use(corsMiddleware())
	r.Use(rateLimit(viper.GetInt64("rps_limit")))
	return setupAuth(r)
}
