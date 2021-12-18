package handlers

// Handlers for exam endpoints, refer to easy_testing.yaml(OpenAPI file) for details.

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gonearewe/EasyTesting/dao"
	"github.com/gonearewe/EasyTesting/models"
	"github.com/gonearewe/EasyTesting/utils"
)

func GetExamHandler(c *gin.Context) {
	teacherId := c.Query("publisher_teacher_id")
	pageSize := utils.Int(c.Query("page_size"))
	pageIndex := utils.Int(c.Query("page_index"))
	exams,num := dao.GetExamsBy(teacherId, pageSize, pageIndex)
	c.JSON(200, gin.H{"total":num, "data": exams})
}

func PostExamHandler(c *gin.Context) {
	abortIfAnyExamActive(c)
	var exams []*models.Exam
	utils.MustParseJsonTo(c, exams)
	dao.CreateExams(exams)
}

func PutExamHandler(c *gin.Context) {
	abortIfAnyExamActive(c)
	var exams []*models.Exam
	utils.MustParseJsonTo(c, exams)
	dao.UpdateExams(exams)
}

func DeleteExamHandler(c *gin.Context) {
	abortIfAnyExamActive(c)
	li := strings.Split(c.Query("ids"), ",")
	ids := make([]int, len(li))
	for i, e := range li {
		ids[i] = utils.Int(e)
	}
	dao.DeleteExams(ids)
}

// abortIfAnyExamActive aborts current request chain if any exam is active,
// this usually happens when trying POST, PUT or DELETE exam-related items (such as questions).
func abortIfAnyExamActive(c *gin.Context){
	if isAnyExamActive(){
		c.AbortWithStatus(http.StatusForbidden)
	}
}

func isAnyExamActive() bool  {
	exams := dao.GetExams()
	for _, exam := range exams{
		if exam.StartTime.Before(time.Now()) && exam.EndTime.After(time.Now()){
			return true
		}
	}
	return false
}