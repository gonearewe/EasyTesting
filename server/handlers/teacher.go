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
    teacher := dao.GetTeachersBy(c.Query("teacher_id"), c.Query("name"),
        utils.Int(c.Query("page_size")), utils.Int(c.Query("page_index")))
    c.JSON(200, teacher)
}

func UpdateTeacherHandler(c *gin.Context) {
    var teacher models.Teacher
    utils.MustParseJsonTo(c, &teacher)
    if teacher.Password != "" {
        salt := utils.GenerateSalt()
        teacher.Salt = string(salt)
        teacher.Password = string(utils.CalculatePasswordWithSalt(teacher.Password, salt))
    }
    dao.UpdateTeacherByTeacherId(&teacher)
}

func DeleteTeachersHandler(c *gin.Context) {
    var teacherIds []string
    utils.MustParseJsonTo(c, teacherIds)
    for _, id := range teacherIds {
        dao.DeleteTeacherByTeacherId(id)
    }
}
