package dao

import (
	"errors"
	"time"

	"github.com/gonearewe/EasyTesting/models"
	"github.com/gonearewe/EasyTesting/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetExamSessionsBy(examId int) (ret []*models.ExamSession) {
	// `answer_sheet` field is excluded.
	err := db.Select("student_id", "student_name", "start_time", "end_time", "score").
		Where("exam_id = ?", examId).Find(&ret).Error
	utils.PanicWhen(err)
	return
}

func GetExamSessionBy(studentId string, examId int) (err error, ret *models.ExamSession) {
	err = db.Select("id", "student_name").
		Where("exam_id = ? AND student_id = ?", examId, studentId).First(&ret).Error
	return
}

func GetEndedExamSessionsBy(examId int) (ret []*models.ExamSession) {
	// `answer_sheet` field is excluded.
	db.Select("student_id", "student_name", "start_time", "end_time", "score").
		Where("exam_id = ?", examId).Find(&ret)
	return
}

func EnterExam(studentId string, examId int) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var exam = models.Exam{}
		var err error

		// SELECT FOR SHARE, make sure the exam exists and active
		err = tx.Clauses(clause.Locking{Strength: "SHARE"}).
			Select("id").
			Where("id = ? AND start_time <= CURTIME() AND CURTIME() < end_time", examId).
			First(&exam).Error
		if err != nil {
			return errors.New("specified exam not exists or not active")
		}

		err = tx.Select("id").Where("exam_id = ?", examId).
			Where("student_id = ?", studentId).First(&models.ExamSession{}).Error
		if err == nil {
			return errors.New("specified student already enters this exam")
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		var session = models.ExamSession{
			ExamID:      examId,
			StudentID:   studentId,
			StudentName: GetStudentBy(studentId).Name,
			StartTime:   time.Now(),
			EndTime:     time.Time{},
		}
		err = tx.Create(&session).Error
		if err != nil {
			return err
		}

		{
			var mcqs []models.Mcq
			err = tx.Select("id", "right_answer").Order("RAND()").Limit(int(exam.McqNum)).Find(&mcqs).Error
			if err != nil {
				return err
			}
			for _, mcq := range mcqs {
				err = tx.Create(&models.McqAnswer{
					McqID:         mcq.ID,
					ExamSessionID: session.ID,
					RightAnswer:   mcq.RightAnswer,
				}).Error
				if err != nil {
					return err
				}
			}
		}

		{
			var maqs []models.Maq
			err = tx.Select("id", "right_answer").Order("RAND()").Limit(int(exam.MaqNum)).Find(&maqs).Error
			if err != nil {
				return err
			}
			for _, maq := range maqs {
				err = tx.Create(&models.MaqAnswer{
					MaqID:         maq.ID,
					ExamSessionID: session.ID,
					RightAnswer:   maq.RightAnswer,
				}).Error
				if err != nil {
					return err
				}
			}
		}

		{
			var bfqs []models.Bfq
			err = tx.Select("id").Order("RAND()").Limit(int(exam.BfqNum)).Find(&bfqs).Error
			if err != nil {
				return err
			}
			for _, bfq := range bfqs {
				err = tx.Create(&models.BfqAnswer{
					BfqID:         bfq.ID,
					ExamSessionID: session.ID,
				}).Error
				if err != nil {
					return err
				}
			}
		}

		{
			var tfqs []models.Tfq
			err = tx.Select("id", "answer").Order("RAND()").Limit(int(exam.TfqNum)).Find(&tfqs).Error
			if err != nil {
				return err
			}
			for _, tfq := range tfqs {
				err = tx.Create(&models.TfqAnswer{
					TfqID:         tfq.ID,
					ExamSessionID: session.ID,
					RightAnswer:   tfq.Answer,
				}).Error
				if err != nil {
					return err
				}
			}
		}

		{
			var crqs []models.Crq
			err = tx.Select("id").Order("RAND()").Limit(int(exam.CrqNum)).Find(&crqs).Error
			if err != nil {
				return err
			}
			for _, crq := range crqs {
				err = tx.Create(&models.CrqAnswer{
					CrqID:         crq.ID,
					ExamSessionID: session.ID,
				}).Error
				if err != nil {
					return err
				}
			}
		}

		{
			var cqs []models.Cq
			err = tx.Select("id").Order("RAND()").Limit(int(exam.CqNum)).Find(&cqs).Error
			if err != nil {
				return err
			}
			for _, cq := range cqs {
				err = tx.Create(&models.CqAnswer{
					CqID:         cq.ID,
					ExamSessionID: session.ID,
				}).Error
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
	utils.PanicWhen(err)
}

func GetMyQuestions(examSessionId int) map[string]interface{} {
	var m = map[string]interface{}{}
	err := db.Transaction(func(tx *gorm.DB) error {
		var err error

		{
			var mcqAnswers []*models.McqAnswer
			err = tx.Select("mcq_id").Where("exam_session_id = ?", examSessionId).Find(&mcqAnswers).Error
			if err != nil {
				return err
			}
			var mcqs []*models.Mcq
			var mcqIds = make([]int, len(mcqAnswers))
			for i, a := range mcqAnswers {
				mcqIds[i] = a.McqID
			}
			err = tx.Where("id IN ?", mcqIds).Find(&mcqs).Error
			if err != nil {
				return err
			}
			m["mcq"] = mcqs
		}

		{
			var maqAnswers []*models.MaqAnswer
			err = tx.Select("maq_id").Where("exam_session_id = ?", examSessionId).Find(&maqAnswers).Error
			if err != nil {
				return err
			}
			var maqs []*models.Maq
			var maqIds = make([]int, len(maqAnswers))
			for i, a := range maqAnswers {
				maqIds[i] = a.MaqID
			}
			err = tx.Where("id IN ?", maqIds).Find(&maqs).Error
			if err != nil {
				return err
			}
			m["maq"] = maqs
		}

		{
			var bfqAnswers []*models.BfqAnswer
			err = tx.Select("bfq_id").Where("exam_session_id = ?", examSessionId).Find(&bfqAnswers).Error
			if err != nil {
				return err
			}
			var bfqs []*models.Bfq
			var bfqIds = make([]int, len(bfqAnswers))
			for i, a := range bfqAnswers {
				bfqIds[i] = a.BfqID
			}
			err = tx.Where("id IN ?", bfqIds).Find(&bfqs).Error
			if err != nil {
				return err
			}
			m["bfq"] = bfqs
		}

		{
			var tfqAnswers []*models.TfqAnswer
			err = tx.Select("tfq_id").Where("exam_session_id = ?", examSessionId).Find(&tfqAnswers).Error
			if err != nil {
				return err
			}
			var tfqs []*models.Tfq
			var tfqIds = make([]int, len(tfqAnswers))
			for i, a := range tfqAnswers {
				tfqIds[i] = a.TfqID
			}
			err = tx.Where("id IN ?", tfqIds).Find(&tfqs).Error
			if err != nil {
				return err
			}
			m["tfq"] = tfqs
		}

		{
			var crqAnswers []*models.CrqAnswer
			err = tx.Select("crq_id").Where("exam_session_id = ?", examSessionId).Find(&crqAnswers).Error
			if err != nil {
				return err
			}
			var crqs []*models.Crq
			var crqIds = make([]int, len(crqAnswers))
			for i, a := range crqAnswers {
				crqIds[i] = a.CrqID
			}
			err = tx.Where("id IN ?", crqIds).Find(&crqs).Error
			if err != nil {
				return err
			}
			m["crq"] = crqs
		}

		{
			// var cqAnswers []*models.CqAnswer
			// err = tx.Select("cq_id").Where("exam_session_id = ?", examSessionId).Find(&cqAnswers).Error
			// if err != nil {
			// 	return err
			// }
			var cqs []*models.Cq
			// var cqIds = make([]int, len(cqAnswers))
			// for i, a := range cqAnswers {
			// 	cqIds[i] = a.CqID
			// }
			err = tx.Where("id = 1").Find(&cqs).Error
			if err != nil {
				return err
			}
			m["cq"] = cqs
		}
		return nil
	})
	utils.PanicWhen(err)
	return m
}
