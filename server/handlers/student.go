package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/gonearewe/EasyTesting/dao"
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
