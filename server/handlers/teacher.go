package handlers

import (
    "net/http"
    "strings"

    jwt "github.com/appleboy/gin-jwt/v2"
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
    teachers := dao.GetTeachersBy(c.Query("teacher_id"), c.Query("name"),
        utils.Int(c.Query("page_size")), utils.Int(c.Query("page_index")))
    c.JSON(200, teachers)
}

func GetTeacherNumHandler(c *gin.Context) {
    num := dao.GetTeacherNumBy(c.Query("teacher_id"), c.Query("name"))
    c.JSON(200, num)
}

func PutTeacherHandler(c *gin.Context) {
    var teacher models.Teacher
    utils.MustParseJsonTo(c, &teacher)
    if teacher.Password != "" {
        salt := utils.GenerateSalt()
        teacher.Salt = string(salt)
        teacher.Password = string(utils.CalculatePasswordWithSalt(teacher.Password, salt))
    }
    dao.UpdateTeacherById(&teacher)
}

func DeleteTeacherHandler(c *gin.Context) {
    abortIfAnyExamActive(c)
    li := strings.Split(c.Query("ids"), ",")
    ids := make([]int, len(li))
    for i, e := range li {
        ids[i] = utils.Int(e)
        if jwt.ExtractClaims(c)["id"] == ids[i] { // not allowed to delete himself or herself
            c.AbortWithStatus(http.StatusForbidden)
        }
    }
    dao.DeleteTeachers(ids)
}
