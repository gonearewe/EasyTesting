package handlers

// Handlers for blank filling question (crq) endpoints, refer to easy_testing.yaml(OpenAPI file) for details.

import (
	"net/http"
	"strings"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gopkg.in/errgo.v2/errors"

	"github.com/gonearewe/EasyTesting/dao"
	"github.com/gonearewe/EasyTesting/models"
	"github.com/gonearewe/EasyTesting/utils"
)

func GetCrqHandler(c *gin.Context) {
	teacherId := c.Query("publisher_teacher_id")
	pageSize := utils.Int(c.Query("page_size"))
	pageIndex := utils.Int(c.Query("page_index"))
	crqs, num := dao.GetCrqsBy(teacherId, pageSize, pageIndex)
	res := make([]gin.H, len(crqs))
	for i, crq := range crqs {
		var answers []string
		for _, answer := range []string{crq.Answer1, crq.Answer2, crq.Answer3, crq.Answer4, crq.Answer5, crq.Answer6} {
			if answer != "" {
				answers = append(answers, answer)
			}
		}

		res[i] = gin.H{
			"id": crq.ID,
			"publisher_teacher_id":  crq.PublisherTeacherID,
			"stem":                  crq.Stem,
			"right_answers":         answers,
			"overall_score":         crq.OverallScore,
			"overall_correct_score": crq.OverallCorrectScore,
		}
	}
	c.JSON(200, gin.H{"total": num, "data": res})
}

func PostCrqHandler(c *gin.Context) {

	var reqs []gin.H
	utils.MustParseJsonTo(c, &reqs)
	var crqs = make([]*models.Crq, len(reqs))
	for i, req := range reqs {
		rightAnswers := req["right_answers"].([]interface{})
		for _, answer := range rightAnswers {
			if answer == "" {
				c.AbortWithError(http.StatusBadRequest, errors.New("empty answer text"))
				return
			}
		}

		crqs[i] = &models.Crq{
			PublisherTeacherID: jwt.ExtractClaims(c)["teacher_id"].(string),
			Stem:               req["stem"].(string),
			Answer1:            rightAnswers[0].(string),
			Answer2:            rightAnswers[1].(string),
		}
		if len(rightAnswers) > 2 {
			crqs[i].Answer3 = rightAnswers[2].(string)
		}
		if len(rightAnswers) > 3 {
			crqs[i].Answer4 = rightAnswers[3].(string)
		}
		if len(rightAnswers) > 4 {
			crqs[i].Answer5 = rightAnswers[4].(string)
		}
		if len(rightAnswers) > 5 {
			crqs[i].Answer6 = rightAnswers[5].(string)
		}
	}
	dao.CreateCrqs(crqs)
}

func PutCrqHandler(c *gin.Context) {

	req := utils.MustParseJson(c)
	rightAnswers := req["right_answers"].([]interface{})
	for _, answer := range rightAnswers {
		if answer == "" {
			c.AbortWithError(http.StatusBadRequest, errors.New("empty answer text"))
			return
		}
	}

	crq := &models.Crq{
		ID:      int(req["id"].(float64)),
		Stem:    req["stem"].(string),
		Answer1: rightAnswers[0].(string),
		Answer2: rightAnswers[1].(string),
	}
	if len(rightAnswers) > 2 {
		crq.Answer3 = rightAnswers[2].(string)
	}
	if len(rightAnswers) > 3 {
		crq.Answer4 = rightAnswers[3].(string)
	}
	if len(rightAnswers) > 4 {
		crq.Answer5 = rightAnswers[4].(string)
	}
	if len(rightAnswers) > 5 {
		crq.Answer6 = rightAnswers[5].(string)
	}
	dao.UpdateCrqById(crq)
}

func DeleteCrqHandler(c *gin.Context) {

	li := strings.Split(c.Query("ids"), ",")
	ids := make([]int, len(li))
	for i, e := range li {
		ids[i] = utils.Int(e)
	}
	dao.DeleteCrqs(ids)
}
