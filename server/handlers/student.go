package handlers

import (
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/gonearewe/EasyTesting/dao"
    "github.com/gonearewe/EasyTesting/models"
    "github.com/gonearewe/EasyTesting/utils"
)

func GetStudentsHandler(c *gin.Context) {
    students, num := dao.GetStudentsBy(c.Query("student_id"), c.Query("name"), c.Query("class_id"),
        utils.Int(c.Query("page_size")), utils.Int(c.Query("page_index")))
    c.JSON(200, gin.H{"total": num, "data": students})
}

func StudentsRegisterHandler(c *gin.Context) {
    abortIfAnyExamActiveOrScoreNotCalculated(c)
    var students []*models.Student
    utils.MustParseJsonTo(c, &students)
    dao.CreateStudents(students)
}

func PutStudentHandler(c *gin.Context) {
    abortIfAnyExamActiveOrScoreNotCalculated(c)
    var student models.Student
    utils.MustParseJsonTo(c, &student)
    dao.UpdateStudentById(&student)
}

func DeleteStudentsHandler(c *gin.Context) {
    abortIfAnyExamActiveOrScoreNotCalculated(c)
    li := strings.Split(c.Query("ids"), ",")
    ids := make([]int, len(li))
    for i, e := range li {
        ids[i] = utils.Int(e)
    }
    dao.DeleteStudents(ids)
}
