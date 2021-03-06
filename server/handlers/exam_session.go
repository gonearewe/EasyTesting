package handlers

import (
	"strconv"
	"strings"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"github.com/gonearewe/EasyTesting/dao"
	"github.com/gonearewe/EasyTesting/models"
	"github.com/gonearewe/EasyTesting/utils"
)

func GetExamineeHandler(c *gin.Context) {
	examId := utils.Int(c.Query("exam_id"))
	if dao.IsExamEndedAndScoresNotCalculated(examId) {
		dao.CalculateScores(examId)
	}
	sessions := dao.GetExamSessionsBy(examId)
	c.JSON(200, sessions)
}

func GetExamSessionHandler(c *gin.Context) {
	studentId := c.Query("student_id")
	examId := utils.Int(c.Query("exam_id"))
	err, sessions := dao.GetExamSessionBy(studentId, examId)
	utils.PanicWhen(err)
	c.JSON(200, sessions)
}

func GetMyQuestionsHandler(c *gin.Context) {
	id := int(jwt.ExtractClaims(c)["exam_session_id"].(float64))
	m := dao.GetMyQuestions(id)

	mcqs := m["mcq"].([]*models.Mcq)
	var mcqMaps = make([]map[string]interface{}, len(mcqs))
	for i, mcq := range mcqs {
		mcqMaps[i] = map[string]interface{}{
			"id":      mcq.ID,
			"stem":    mcq.Stem,
			"choices": []string{mcq.Choice1, mcq.Choice2, mcq.Choice3, mcq.Choice4},
		}
	}

	maqs := m["maq"].([]*models.Maq)
	var maqMaps = make([]map[string]interface{}, len(maqs))
	for i, maq := range maqs {
		choices := []string{maq.Choice1, maq.Choice2, maq.Choice3, maq.Choice4}
		if maq.Choice5 != "" {
			choices = append(choices, maq.Choice5)
		}
		if maq.Choice6 != "" {
			choices = append(choices, maq.Choice6)
		}
		if maq.Choice7 != "" {
			choices = append(choices, maq.Choice7)
		}
		maqMaps[i] = map[string]interface{}{
			"id":      maq.ID,
			"stem":    maq.Stem,
			"choices": choices,
		}
	}

	bfqs := m["bfq"].([]*models.Bfq)
	var bfqMaps = make([]map[string]interface{}, len(bfqs))
	for i, bfq := range bfqs {
		bfqMaps[i] = map[string]interface{}{
			"id":        bfq.ID,
			"stem":      bfq.Stem,
			"blank_num": bfq.BlankNum,
		}
	}

	tfqs := m["tfq"].([]*models.Tfq)
	var tfqMaps = make([]map[string]interface{}, len(tfqs))
	for i, tfq := range tfqs {
		tfqMaps[i] = map[string]interface{}{
			"id":   tfq.ID,
			"stem": tfq.Stem,
		}
	}

	crqs := m["crq"].([]*models.Crq)
	var crqMaps = make([]map[string]interface{}, len(crqs))
	for i, crq := range crqs {
		crqMaps[i] = map[string]interface{}{
			"id":        crq.ID,
			"stem":      crq.Stem,
			"blank_num": crq.BlankNum,
		}
	}

	cqs := m["cq"].([]*models.Cq)
	var cqMaps = make([]map[string]interface{}, len(cqs))
	for i, cq := range cqs {
		cqMaps[i] = map[string]interface{}{
			"id":                 cq.ID,
			"stem":               cq.Stem,
			"is_input_from_file": cq.IsInputFromFile,
			"is_output_to_file":  cq.IsOutputToFile,
			"template":           cq.Template,
			"input":              cq.Input,
		}
	}

	c.JSON(200, gin.H{
		"mcq": mcqMaps,
		"maq": maqMaps,
		"bfq": bfqMaps,
		"tfq": tfqMaps,
		"crq": crqMaps,
		"cq":  cqMaps,
	})
}

func PutMyAnswersHandler(c *gin.Context) {
	var myAnswers models.MyAnswers
	utils.MustParseJsonTo(c, &myAnswers)
	task := newTask()

	for _, m := range myAnswers.Mcq {
		if m["answer"] == nil {
			continue
		}
		task.mcqs = append(task.mcqs, &models.McqAnswer{
			McqID:         int(m["id"].(float64)),
			StudentAnswer: strconv.Itoa(int(m["answer"].(float64)))})
	}

	// if maq answer is none(no choice selected), then it won't be recorded
	for _, m := range myAnswers.Maq {
		answer := m["answer"].([]interface{})
		if len(answer) == 0 {
			continue
		}
		var tmp = make([]string, len(answer))
		for i, e := range answer {
			tmp[i] = strconv.Itoa(int(e.(float64)))
		}
		task.maqs = append(task.maqs, &models.MaqAnswer{
			MaqID:         int(m["id"].(float64)),
			StudentAnswer: strings.Join(tmp, ""),
		})
	}

	for _, m := range myAnswers.Bfq {
		answer := m["answer"].([]interface{})
		tmp := &models.BfqAnswer{
			BfqID: int(m["id"].(float64)),
		}
		if answer[0] != nil {
			tmp.StudentAnswer1 = answer[0].(string)
		}
		if len(answer) > 1 && answer[1] != nil {
			tmp.StudentAnswer2 = answer[1].(string)
		}
		if len(answer) > 2 && answer[2] != nil {
			tmp.StudentAnswer3 = answer[2].(string)
		}
		task.bfqs = append(task.bfqs, tmp)
	}

	for _, m := range myAnswers.Tfq {
		if m["answer"] == nil {
			continue
		}
		task.tfqs = append(task.tfqs, &models.TfqAnswer{
			TfqID:         int(m["id"].(float64)),
			StudentAnswer: m["answer"].(bool),
		})
	}

	for _, m := range myAnswers.Crq {
		answer := m["answer"].([]interface{})
		tmp := &models.CrqAnswer{
			CrqID: int(m["id"].(float64)),
		}
		if answer[0] != nil {
			tmp.StudentAnswer1 = answer[0].(string)
		}
		if len(answer) > 1 && answer[1] != nil {
			tmp.StudentAnswer2 = answer[1].(string)
		}
		if len(answer) > 2 && answer[2] != nil {
			tmp.StudentAnswer3 = answer[2].(string)
		}
		if len(answer) > 3 && answer[3] != nil {
			tmp.StudentAnswer4 = answer[3].(string)
		}
		if len(answer) > 4 && answer[4] != nil {
			tmp.StudentAnswer5 = answer[4].(string)
		}
		if len(answer) > 5 && answer[5] != nil {
			tmp.StudentAnswer6 = answer[5].(string)
		}
		task.crqs = append(task.crqs, tmp)
	}

	for _, m := range myAnswers.Cq {
		task.cqs = append(task.cqs, &models.CqAnswer{
			CqID:          int(m["id"].(float64)),
			StudentAnswer: m["answer"].(string),
			StudentOutput: m["output"].(string),
		})
	}

	task.examSessionId = int(jwt.ExtractClaims(c)["exam_session_id"].(float64))
	taskQueue <- task
}
