package handlers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	var id string
	var ok bool
	if id, ok = jwt.ExtractClaims(c)["student_id"].(string); !ok {
		id = jwt.ExtractClaims(c)["teacher_id"].(string)
	}
	c.JSON(200, gin.H{
		"greeting": "hello",
		"id":       id,
	})
}
