package handlers

// Handlers for blank filling question (bfq) endpoints, refer to easy_testing.yaml(OpenAPI file) for details.

import (
    "net/http"
    "strings"

    jwt "github.com/appleboy/gin-jwt/v2"
    "github.com/gin-gonic/gin"
    "github.com/gonearewe/EasyTesting/dao"
    "github.com/gonearewe/EasyTesting/models"
    "github.com/gonearewe/EasyTesting/utils"
    "gopkg.in/errgo.v2/errors"
)

func GetBfqHandler(c *gin.Context) {
    teacherId := c.Query("publisher_teacher_id")
    pageSize := utils.Int(c.Query("page_size"))
    pageIndex := utils.Int(c.Query("page_index"))
    bfqs, num := dao.GetBfqsBy(teacherId, pageSize, pageIndex)
    res := make([]gin.H, len(bfqs))
    for i, bfq := range bfqs {
        var answers []string
        for _, answer := range []string{bfq.Answer1, bfq.Answer2, bfq.Answer3} {
            if answer != "" {
                answers = append(answers, answer)
            }
        }

        res[i] = gin.H{
            "id":                   bfq.ID,
            "publisher_teacher_id": bfq.PublisherTeacherID,
            "stem":                 bfq.Stem,
            "right_answers":        answers,
        }
    }
    c.JSON(200, gin.H{"total": num, "data": res})
}

func PostBfqHandler(c *gin.Context) {
    abortIfAnyExamActive(c)
    var reqs []gin.H
    utils.MustParseJsonTo(c, &reqs)
    var bfqs = make([]*models.Bfq, len(reqs))
    for i, req := range reqs {
        rightAnswers := req["right_answers"].([]interface{})
        if len(rightAnswers) > 3 {
            c.AbortWithError(http.StatusBadRequest, errors.New("length of right_answers more than 3"))
            return
        }
        for _, answer := range rightAnswers {
            if answer == "" {
                c.AbortWithError(http.StatusBadRequest, errors.New("empty answer text"))
                return
            }
        }

        bfqs[i] = &models.Bfq{
            PublisherTeacherID: jwt.ExtractClaims(c)["teacher_id"].(string),
            Stem:               req["stem"].(string),
            Answer1:            rightAnswers[0].(string),
        }
        if len(rightAnswers) > 1 {
            bfqs[i].Answer2 = rightAnswers[1].(string)
        }
        if len(rightAnswers) > 2 {
            bfqs[i].Answer3 = rightAnswers[2].(string)
        }
    }
    dao.CreateBfqs(bfqs)
}

func PutBfqHandler(c *gin.Context) {
    abortIfAnyExamActive(c)
    req := utils.MustParseJson(c)
    rightAnswers := req["right_answers"].([]interface{})
    if len(rightAnswers) > 3 {
        c.AbortWithError(http.StatusBadRequest, errors.New("length of right_answers more than 3"))
        return
    }
    for _, answer := range rightAnswers {
        if answer == "" {
            c.AbortWithError(http.StatusBadRequest, errors.New("empty answer text"))
            return
        }
    }

    bfq := &models.Bfq{
        ID:      int(req["id"].(float64)),
        Stem:    req["stem"].(string),
        Answer1: rightAnswers[0].(string),
    }
    if len(rightAnswers) > 1 {
        bfq.Answer2 = rightAnswers[1].(string)
    }
    if len(rightAnswers) > 2 {
        bfq.Answer3 = rightAnswers[2].(string)
    }
    dao.UpdateBfqById(bfq)
}

func DeleteBfqHandler(c *gin.Context) {
    abortIfAnyExamActive(c)
    li := strings.Split(c.Query("ids"), ",")
    ids := make([]int, len(li))
    for i, e := range li {
        ids[i] = utils.Int(e)
    }
    dao.DeleteBfqs(ids)
}
