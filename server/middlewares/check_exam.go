package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gonearewe/EasyTesting/dao"

	"gopkg.in/errgo.v2/errors"
)

func CheckGlobalExamStatus(c *gin.Context) {
	// aborts current request chain if any exam is active
	// or scores of who participated it aren't calculated,
	// this usually happens when trying POST, PUT or DELETE exam-related items (such as questions).
	if dao.AnyExamActiveOrScoreNotCalculated() {
		c.AbortWithError(http.StatusForbidden,
			errors.New("action forbidden when any exam is active or its scores aren't calculated"))
	} else {
		c.Next()
	}
}
