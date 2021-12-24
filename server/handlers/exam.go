package handlers

// Handlers for exam endpoints, refer to easy_testing.yaml(OpenAPI file) for details.

import (
    "net/http"
    "strings"
    "time"

    jwt "github.com/appleboy/gin-jwt/v2"
    "github.com/gin-gonic/gin"
    "github.com/gonearewe/EasyTesting/dao"
    "github.com/gonearewe/EasyTesting/models"
    "github.com/gonearewe/EasyTesting/utils"
    "gopkg.in/errgo.v2/errors"
)

func GetExamHandler(c *gin.Context) {
    teacherId := c.Query("publisher_teacher_id")
    pageSize := utils.Int(c.Query("page_size"))
    pageIndex := utils.Int(c.Query("page_index"))
    exams, num := dao.GetExamsBy(teacherId, pageSize, pageIndex)
    c.JSON(200, gin.H{"total": num, "data": exams})
}

func GetEndedExamHandler(c *gin.Context) {
    exams := dao.GetEndedExams()
    c.JSON(200, exams)
}

func PostExamHandler(c *gin.Context) {
    var exams []*models.Exam
    utils.MustParseJsonTo(c, &exams)
    validateExams(exams)
    for _, exam := range exams {
        exam.PublisherTeacherID = jwt.ExtractClaims(c)["teacher_id"].(string)
    }
    dao.CreateExams(exams)
}

func PutExamHandler(c *gin.Context) {
    var exams []*models.Exam
    utils.MustParseJsonTo(c, &exams)
    validateExams(exams)
    dao.UpdateExams(exams)
}

func validateExams(exams []*models.Exam) {
    for _, exam := range exams {
        if exam.StartTime.Before(time.Now()) {
            utils.PanicWhen(errors.New("Exam with passed `StartTime` is forbidden"))
        } else if exam.EndTime.Sub(exam.StartTime).Milliseconds() <= (int64(exam.TimeAllowed)+10)*60*1000 {
            utils.PanicWhen(
                errors.New("Exam with `EndTime - StartTime` no more than `TimeAllowed + 10 minutes` is forbidden"))
        }
    }
}

func DeleteExamHandler(c *gin.Context) {
    li := strings.Split(c.Query("ids"), ",")
    ids := make([]int, len(li))
    for i, e := range li {
        ids[i] = utils.Int(e)
    }
    dao.DeleteExams(ids)
}

// abortIfAnyExamActive aborts current request chain if any exam is active,
// this usually happens when trying POST, PUT or DELETE exam-related items (such as questions).
func abortIfAnyExamActive(c *gin.Context) {
    if isAnyExamActive() {
        c.AbortWithStatus(http.StatusForbidden)
    }
}

func isAnyExamActive() bool {
    exams := dao.GetExams()
    for _, exam := range exams {
        if exam.StartTime.Before(time.Now()) && exam.EndTime.After(time.Now()) {
            return true
        }
    }
    return false
}


