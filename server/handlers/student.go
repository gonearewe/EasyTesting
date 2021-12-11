package handlers

import (
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/gonearewe/EasyTesting/dao"
    "github.com/gonearewe/EasyTesting/models"
    "github.com/gonearewe/EasyTesting/utils"
)

func GetStudentsHandler(c *gin.Context) {
    students := dao.GetStudentsBy(c.Query("student_id"), c.Query("name"),c.Query("class_id"),
        utils.Int(c.Query("page_size")), utils.Int(c.Query("page_index")))
    c.JSON(200, students)
}

func GetStudentNumHandler(c *gin.Context) {
    num := dao.GetStudentNumBy(c.Query("student_id"), c.Query("name"),c.Query("class_id"))
    c.JSON(200, num)
}

func StudentsRegisterHandler(c *gin.Context) {
    var students []*models.Student
    utils.MustParseJsonTo(c, &students)
    dao.CreateStudents(students)
}

func PutStudentHandler(c *gin.Context) {
    var student models.Student
    utils.MustParseJsonTo(c, &student)
    dao.UpdateStudentById(&student)
}

func DeleteStudentsHandler(c *gin.Context) {
    abortIfAnyExamActive(c)
    li := strings.Split(c.Query("ids"), ",")
    ids := make([]int, len(li))
    for i, e := range li {
        ids[i] = utils.Int(e)
    }
    dao.DeleteStudents(ids)
}
