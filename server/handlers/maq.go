package handlers

// Handlers for multiple answer question (maq) endpoints, refer to easy_testing.yaml(OpenAPI file) for details.

import (
    "strings"

    jwt "github.com/appleboy/gin-jwt/v2"
    "github.com/gin-gonic/gin"
    "github.com/gonearewe/EasyTesting/dao"
    "github.com/gonearewe/EasyTesting/models"
    "github.com/gonearewe/EasyTesting/utils"
)

func GetMaqHandler(c *gin.Context) {
    teacherId := c.Query("publisher_teacher_id")
    pageSize := utils.Int(c.Query("page_size"))
    pageIndex := utils.Int(c.Query("page_index"))
    maqs := dao.GetMaqsBy(teacherId, pageSize, pageIndex)
    res := make([]gin.H, len(maqs))
    for i, maq := range maqs {
        var choices []string
        for _, choice := range []string{
            maq.Choice1, maq.Choice2, maq.Choice3, maq.Choice4, maq.Choice5, maq.Choice6, maq.Choice7} {
            if choice != "" {
                choices = append(choices, choice)
            }
        }

        res[i] = gin.H{
            "id":                   maq.ID,
            "publisher_teacher_id": maq.PublisherTeacherID,
            "stem":                 maq.Stem,
            "choices":              choices,
            "right_answer":         maq.RightAnswer,
        }
    }
    c.JSON(200, res)
}

func PostMaqHandler(c *gin.Context) {
    abortIfAnyExamActive(c)
    var reqs []gin.H
    utils.MustParseJsonTo(c, &reqs)
    var maqs = make([]*models.Maq, len(reqs))
    for i, req := range reqs {
        choices := req["choices"].([]interface{})
        maqs[i] = &models.Maq{
            PublisherTeacherID: jwt.ExtractClaims(c)["teacher_id"].(string),
            Stem:               req["stem"].(string),
            Choice1:            choices[0].(string),
            Choice2:            choices[1].(string),
            Choice3:            choices[2].(string),
            Choice4:            choices[3].(string),
            RightAnswer:        req["right_answer"].(string),
        }
        if len(choices) > 4 {
            maqs[i].Choice4 = choices[4].(string)
        }
        if len(choices) > 5 {
            maqs[i].Choice5 = choices[5].(string)
        }
        if len(choices) > 6 {
            maqs[i].Choice6 = choices[6].(string)
        }
    }
    dao.CreateMaqs(maqs)
}

func PutMaqHandler(c *gin.Context) {
    abortIfAnyExamActive(c)
    var reqs []gin.H
    utils.MustParseJsonTo(c, &reqs)
    var maqs = make([]*models.Maq, len(reqs))
    for i, req := range reqs {
        choices := req["choices"].([]interface{})
        maqs[i] = &models.Maq{
            ID:          int(req["id"].(float64)),
            Stem:        req["stem"].(string),
            Choice1:     choices[0].(string),
            Choice2:     choices[1].(string),
            Choice3:     choices[2].(string),
            Choice4:     choices[3].(string),
            RightAnswer: req["right_answer"].(string),
        }
        if len(choices) > 4 {
            maqs[i].Choice4 = choices[4].(string)
        }
        if len(choices) > 5 {
            maqs[i].Choice5 = choices[5].(string)
        }
        if len(choices) > 6 {
            maqs[i].Choice6 = choices[6].(string)
        }
    }
    dao.UpdateMaqs(maqs)
}

func DeleteMaqHandler(c *gin.Context) {
    abortIfAnyExamActive(c)
    li := strings.Split(c.Query("ids"), ",")
    ids := make([]int, len(li))
    for i, e := range li {
        ids[i] = utils.Int(e)
    }
    dao.DeleteMaqs(ids)
}
