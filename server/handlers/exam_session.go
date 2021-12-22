package handlers

import (
    "github.com/gin-gonic/gin"
    "github.com/gonearewe/EasyTesting/dao"
    "github.com/gonearewe/EasyTesting/utils"
)

func EnterExamHandler(c *gin.Context) {
    m := utils.MustParseJson(c)
    dao.EnterExam(m["student_id"].(string), int(m["exam_id"].(float64)))
}

func GetExamineeHandler(c *gin.Context){
    examId := utils.Int(c.Query("exam_id"))
    sessions := dao.GetExamSessionsBy(examId)
    c.JSON(200,sessions)
}
