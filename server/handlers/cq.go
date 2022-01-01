package handlers

// Handlers for blank filling question (cq) endpoints, refer to easy_testing.yaml(OpenAPI file) for details.

import (
    "strings"

    jwt "github.com/appleboy/gin-jwt/v2"
    "github.com/gin-gonic/gin"
    "github.com/gonearewe/EasyTesting/dao"
    "github.com/gonearewe/EasyTesting/models"
    "github.com/gonearewe/EasyTesting/utils"
)

func GetCqHandler(c *gin.Context) {
    teacherId := c.Query("publisher_teacher_id")
    pageSize := utils.Int(c.Query("page_size"))
    pageIndex := utils.Int(c.Query("page_index"))
    cqs, num := dao.GetCqsBy(teacherId, pageSize, pageIndex)
    c.JSON(200, gin.H{"total": num, "data": cqs})
}

func PostCqHandler(c *gin.Context) {
    abortIfAnyExamActive(c)
    var cqs []*models.Cq
    utils.MustParseJsonTo(c, &cqs)
    for _,cq:= range cqs{
        cq.PublisherTeacherID = jwt.ExtractClaims(c)["teacher_id"].(string)
    }
    dao.CreateCqs(cqs)
}

func PutCqHandler(c *gin.Context) {
    abortIfAnyExamActive(c)
    var cq models.Cq
    utils.MustParseJsonTo(c, &cq)
    dao.UpdateCqById(&cq)
}

func DeleteCqHandler(c *gin.Context) {
    abortIfAnyExamActive(c)
    li := strings.Split(c.Query("ids"), ",")
    ids := make([]int, len(li))
    for i, e := range li {
        ids[i] = utils.Int(e)
    }
    dao.DeleteCqs(ids)
}
