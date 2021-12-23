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
    db.Select("student_id","student_name","start_time","end_time","score").
        Where("exam_id = ?", examId).Find(&ret)
    return
}

func GetEndedExamSessionsBy(examId int) (ret []*models.ExamSession) {
    // `answer_sheet` field is excluded.
    db.Select("student_id","student_name","start_time","end_time","score").
        Where("exam_id = ?", examId).Find(&ret)
    return
}

func EnterExam(studentId string, examId int) {
    err := db.Transaction(func(tx *gorm.DB) error {
        var exam =models.Exam{}
        var err error

        // SELECT FOR SHARE, make sure the exam exists and active
        err = tx.Clauses(clause.Locking{Strength: "SHARE"}).
            Select("id").
            Where("id = ? AND start_time <= CURTIME() AND CURTIME() < end_time", examId).
            First(&exam).Error
        if err != nil {
            return errors.New("specified exam not exists or not active")
        }

        err = tx.Select("id").Where("exam_id = ?",examId).
            Where("student_id = ?",studentId).First(&models.ExamSession{}).Error
        if err == nil {
           return  errors.New("specified student already enters this exam")
        }else if !errors.Is(err,gorm.ErrRecordNotFound) {
            return err
        }

        var session = models.ExamSession{
            ExamID:      examId,
            StudentID:   studentId,
            StudentName: GetStudentNameBy(studentId),
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
                    RightAnswer: tfq.Answer,
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

        return nil
    })
    utils.PanicWhen(err)
}

// func GetMyQuestions(examSessionId int) {
//
//
// }
