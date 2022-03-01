package middlewares

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/google/logger"
)

func recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			trace := make([]byte, 1<<16)
			n := runtime.Stack(trace, false)
			logger.Errorf("%v\n%s", err, string(trace[:n]))
			// logger.Error(err)
			c.AbortWithStatus(400)
		}
	}()
	c.Next()
}
