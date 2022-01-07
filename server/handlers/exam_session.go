package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gonearewe/EasyTesting/dao"
	"github.com/gonearewe/EasyTesting/models"
	"github.com/gonearewe/EasyTesting/utils"
)

func EnterExamHandler(c *gin.Context) {
	m := utils.MustParseJson(c)
	dao.EnterExam(m["student_id"].(string), int(m["exam_id"].(float64)))
}

func GetExamineeHandler(c *gin.Context) {
	examId := utils.Int(c.Query("exam_id"))
	sessions := dao.GetExamSessionsBy(examId)
	c.JSON(200, sessions)
}

func GetExamSessionHandler(c *gin.Context) {
	studentId := c.Query("student_id")
	examId := utils.Int(c.Query("exam_id"))
	sessions := dao.GetExamSessionBy(studentId, examId)
	c.JSON(200, sessions)
}

func GetMyQuestionsHandler(c *gin.Context) {
	id := utils.Int(c.Query("exam_session_id"))
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
			"output":             cq.Output,
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
