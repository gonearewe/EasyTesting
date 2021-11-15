package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gonearewe/EasyTesting/dao"
	"github.com/gonearewe/EasyTesting/models"
	"github.com/gonearewe/EasyTesting/utils"
)

func TeachersRegisterHandler(c *gin.Context) {
	var teachers []*models.Teacher
	utils.MustParseJsonTo(c, &teachers)
	for _, teacher := range teachers {
		salt := utils.GenerateSalt()
		teacher.Salt = string(salt)
		teacher.Password = string(utils.CalculatePasswordWithSalt(teacher.Password, salt))
	}
	dao.CreateTeachers(teachers)
}

func GetTeachersHandler(c *gin.Context) {
	json := utils.MustParseJson(c)
	teacher := dao.GetTeachersBy(json["teacher_id"].(string), json["name"].(string),
		int(json["page_size"].(float64)), int(json["page_index"].(float64)))
	c.JSON(200, teacher)
}
