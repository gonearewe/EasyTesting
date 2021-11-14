package handlers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gonearewe/EasyTesting/dao"
	"github.com/gonearewe/EasyTesting/utils"
)

func TeacherAuthenticator(c *gin.Context) (user interface{}, err error) {
	defer func() {
		if recover() != nil {
			user, err = nil, jwt.ErrFailedAuthentication
		}
	}()
	json := utils.MustParseJson(c)
	id := json["teacher_id"].(string)
	teacher := dao.GetTeacherByTeacherId(id)
	err = utils.VerifyPassword(json["password"].(string), teacher.Salt, teacher.Password)
	utils.PanicWhen(err)
	return teacher, nil
}
