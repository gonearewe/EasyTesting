package handlers

// Handlers for multiple answer question (maq) endpoints, refer to easy_testing.yaml(OpenAPI file) for details.

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

func GetMaqHandler(c *gin.Context) {
	teacherId := c.Query("publisher_teacher_id")
	pageSize := utils.Int(c.Query("page_size"))
	pageIndex := utils.Int(c.Query("page_index"))
	maqs, num := dao.GetMaqsBy(teacherId, pageSize, pageIndex)
	res := make([]gin.H, len(maqs))
	for i, maq := range maqs {
		var choices []string
		for _, choice := range []string{
			maq.Choice1, maq.Choice2, maq.Choice3, maq.Choice4, maq.Choice5, maq.Choice6, maq.Choice7} {
			if choice != "" {
				choices = append(choices, choice)
			}
		}

		li := strings.Split(maq.RightAnswer, "")
		answer := make([]int, len(maq.RightAnswer))
		for i := range answer {
			answer[i] = utils.Int(li[i])
		}
		res[i] = gin.H{
			"id": maq.ID,
			"publisher_teacher_id":  maq.PublisherTeacherID,
			"stem":                  maq.Stem,
			"choices":               choices,
			"right_answer":          answer,
			"overall_score":         maq.OverallScore,
			"overall_correct_score": maq.OverallCorrectScore,
		}
	}
	c.JSON(200, gin.H{"total": num, "data": res})
}

func PostMaqHandler(c *gin.Context) {

	var reqs []gin.H
	utils.MustParseJsonTo(c, &reqs)
	var maqs = make([]*models.Maq, len(reqs))
	for i, req := range reqs {
		choices := req["choices"].([]interface{})
		if len(choices) < 4 {
			c.AbortWithError(http.StatusBadRequest, errors.New("length of choices less than 4"))
			return
		}
		for _, choice := range choices {
			if choice == "" {
				c.AbortWithError(http.StatusBadRequest, errors.New("empty choice text"))
				return
			}
		}

		maqs[i] = &models.Maq{
			PublisherTeacherID: jwt.ExtractClaims(c)["teacher_id"].(string),
			Stem:               req["stem"].(string),
			Choice1:            choices[0].(string),
			Choice2:            choices[1].(string),
			Choice3:            choices[2].(string),
			Choice4:            choices[3].(string),
			RightAnswer:        utils.Join(req["right_answer"].([]interface{})),
		}
		if len(choices) > 4 {
			maqs[i].Choice5 = choices[4].(string)
		}
		if len(choices) > 5 {
			maqs[i].Choice6 = choices[5].(string)
		}
		if len(choices) > 6 {
			maqs[i].Choice7 = choices[6].(string)
		}
	}
	dao.CreateMaqs(maqs)
}

func PutMaqHandler(c *gin.Context) {

	req := utils.MustParseJson(c)
	choices := req["choices"].([]interface{})
	if len(choices) < 4 {
		c.AbortWithError(http.StatusBadRequest, errors.New("length of choices less than 4"))
		return
	}
	for _, choice := range choices {
		if choice == "" {
			c.AbortWithError(http.StatusBadRequest, errors.New("empty choice text"))
			return
		}
	}

	maq := &models.Maq{
		ID:          int(req["id"].(float64)),
		Stem:        req["stem"].(string),
		Choice1:     choices[0].(string),
		Choice2:     choices[1].(string),
		Choice3:     choices[2].(string),
		Choice4:     choices[3].(string),
		RightAnswer: utils.Join(req["right_answer"].([]interface{})),
	}
	if len(choices) > 4 {
		maq.Choice5 = choices[4].(string)
	}
	if len(choices) > 5 {
		maq.Choice6 = choices[5].(string)
	}
	if len(choices) > 6 {
		maq.Choice7 = choices[6].(string)
	}
	dao.UpdateMaqById(maq)
}

func DeleteMaqHandler(c *gin.Context) {

	li := strings.Split(c.Query("ids"), ",")
	ids := make([]int, len(li))
	for i, e := range li {
		ids[i] = utils.Int(e)
	}
	dao.DeleteMaqs(ids)
}
