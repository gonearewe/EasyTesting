package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
)

func recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
// 			logger.Errorf("%v\n", err)
			logger.Error(err)
			c.AbortWithStatus(400)
		}
	}()
	c.Next()
}
