package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func recovery(c *gin.Context) {
	defer func() {
		if e := recover(); e != nil {
			if err, ok := e.(error); ok {
				log.Error().Stack().Err(err).Msg("Gin recover")
			} else {
				log.Error().Stack().Interface("panic value", e).Msg("Gin recover")
			}
		}
	}()
	c.Next()
}
