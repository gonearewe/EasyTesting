package handlers

// Handlers for multiple choice question (mcq) endpoints, refer to easy_testing.yaml(OpenAPI file) for details.

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"github.com/gonearewe/EasyTesting/dao"
	"github.com/gonearewe/EasyTesting/models"
	"github.com/gonearewe/EasyTesting/utils"
)

func GetMcqHandler(c *gin.Context) {
	teacherId := c.Query("publisher_teacher_id")
	pageSize := utils.Int(c.Query("page_size"))
	pageIndex := utils.Int(c.Query("page_index"))
	mcqs, num := dao.GetMcqsBy(teacherId, pageSize, pageIndex)
	res := make([]gin.H, len(mcqs))
	for i, mcq := range mcqs {
		res[i] = gin.H{
			"id": mcq.ID,
			"publisher_teacher_id":  mcq.PublisherTeacherID,
			"stem":                  mcq.Stem,
			"choices":               [4]string{mcq.Choice1, mcq.Choice2, mcq.Choice3, mcq.Choice4},
			"right_answer":          utils.Int(mcq.RightAnswer),
			"overall_score":         mcq.OverallScore,
			"overall_correct_score": mcq.OverallCorrectScore,
		}
	}
	c.JSON(200, gin.H{"total": num, "data": res})
}

func PostMcqHandler(c *gin.Context) {

	var reqs []gin.H
	utils.MustParseJsonTo(c, &reqs)
	var mcqs = make([]*models.Mcq, len(reqs))
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

		mcqs[i] = &models.Mcq{
			PublisherTeacherID: jwt.ExtractClaims(c)["teacher_id"].(string),
			Stem:               req["stem"].(string),
			Choice1:            choices[0].(string),
			Choice2:            choices[1].(string),
			Choice3:            choices[2].(string),
			Choice4:            choices[3].(string),
			RightAnswer:        strconv.Itoa(int(req["right_answer"].(float64))),
		}
	}
	dao.CreateMcqs(mcqs)
}

func PutMcqHandler(c *gin.Context) {

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

	mcq := &models.Mcq{
		ID:          int(req["id"].(float64)),
		Stem:        req["stem"].(string),
		Choice1:     choices[0].(string),
		Choice2:     choices[1].(string),
		Choice3:     choices[2].(string),
		Choice4:     choices[3].(string),
		RightAnswer: strconv.Itoa(int(req["right_answer"].(float64))),
	}
	dao.UpdateMcqById(mcq)
}

func DeleteMcqHandler(c *gin.Context) {

	li := strings.Split(c.Query("ids"), ",")
	ids := make([]int, len(li))
	for i, e := range li {
		ids[i] = utils.Int(e)
	}
	dao.DeleteMcqs(ids)
}
