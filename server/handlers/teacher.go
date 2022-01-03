package handlers

import (
    "strings"

    jwt "github.com/appleboy/gin-jwt/v2"
    "github.com/gin-gonic/gin"
    "github.com/gonearewe/EasyTesting/dao"
    "github.com/gonearewe/EasyTesting/models"
    "github.com/gonearewe/EasyTesting/utils"
    "gopkg.in/errgo.v2/errors"
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
    teachers,num := dao.GetTeachersBy(c.Query("teacher_id"), c.Query("name"),
        utils.Int(c.Query("page_size")), utils.Int(c.Query("page_index")))
    c.JSON(200, gin.H{"total":num, "data":teachers})
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

func PutTeacherProfileHandler(c *gin.Context) {
    var teacher models.Teacher
    utils.MustParseJsonTo(c, &teacher)
    if teacher.ID != int(jwt.ExtractClaims(c)["id"].(float64)){
        c.AbortWithError(401,errors.New("profile endpoint forbids modifying other teacher's profile"))
        return
    }
    if teacher.Password != "" {
        salt := utils.GenerateSalt()
        teacher.Salt = string(salt)
        teacher.Password = string(utils.CalculatePasswordWithSalt(teacher.Password, salt))
    }
    dao.UpdateTeacherProfileById(&teacher)
}

func DeleteTeachersHandler(c *gin.Context) {
    abortIfAnyExamActive(c)
    li := strings.Split(c.Query("ids"), ",")
    ids := make([]int, len(li))
    for i, e := range li {
        ids[i] = utils.Int(e)
        if int(jwt.ExtractClaims(c)["id"].(float64)) == ids[i] {
            c.AbortWithError(403,errors.New("not allowed to delete himself or herself"))
            return
        }
    }
    dao.DeleteTeachers(ids)
}
