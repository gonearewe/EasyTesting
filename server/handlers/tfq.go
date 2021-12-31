package handlers

// Handlers for multiple choice question (tfq) endpoints, refer to easy_testing.yaml(OpenAPI file) for details.

import (
	"strings"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/gonearewe/EasyTesting/dao"
	"github.com/gonearewe/EasyTesting/models"
	"github.com/gonearewe/EasyTesting/utils"
)

func GetTfqHandler(c *gin.Context) {
	teacherId := c.Query("publisher_teacher_id")
	pageSize := utils.Int(c.Query("page_size"))
	pageIndex := utils.Int(c.Query("page_index"))
	tfqs,num := dao.GetTfqsBy(teacherId, pageSize, pageIndex)
	res := make([]gin.H, len(tfqs))
	for i, tfq := range tfqs {
		res[i] = gin.H{
			"id":                   tfq.ID,
			"publisher_teacher_id": tfq.PublisherTeacherID,
			"stem":                 tfq.Stem,
			"right_answer":         tfq.Answer,
		}
	}
    c.JSON(200, gin.H{"total":num,"data":res})
}

func PostTfqHandler(c *gin.Context) {
	abortIfAnyExamActive(c)
	var reqs []gin.H
	utils.MustParseJsonTo(c, &reqs)
	var tfqs = make([]*models.Tfq, len(reqs))
	for i, req := range reqs {
		tfqs[i] = &models.Tfq {
			PublisherTeacherID: jwt.ExtractClaims(c)["teacher_id"].(string),
			Stem:               req["stem"].(string),
			Answer:              req["right_answer"].(bool),
		}
	}
	dao.CreateTfqs(tfqs)
}

func PutTfqHandler(c *gin.Context) {
	abortIfAnyExamActive(c)
	req := utils.MustParseJson(c)
	tfq := &models.Tfq{
		ID:          int(req["id"].(float64)),
		Stem:        req["stem"].(string),
		Answer: req["right_answer"].(bool),
	}
	dao.UpdateTfqById(tfq)
}

func DeleteTfqHandler(c *gin.Context) {
	abortIfAnyExamActive(c)
	li := strings.Split(c.Query("ids"), ",")
	ids := make([]int, len(li))
	for i, e := range li {
		ids[i] = utils.Int(e)
	}
	dao.DeleteTfqs(ids)
}
