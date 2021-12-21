package dao

import (
    "time"

    "github.com/gonearewe/EasyTesting/models"
    "github.com/gonearewe/EasyTesting/utils"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

// GetExamsBy searches the database for exam whose publisher teacher id (string) starts
// with `teacherId`, it only returns records in given `pageIndex` (1-based) in the increasing order of id,
// but plus the total number of all filtered.
// When any error occurs, it panics.
func GetExamsBy(teacherId string, pageSize int, pageIndex int) (ret []*models.Exam, num int64) {
    err := db.Transaction(func(tx *gorm.DB) error {
        err := buildQueryFrom(tx, teacherId, models.Exam{}).
            Limit(pageSize).Offset(pageSize * (pageIndex - 1)).
            Find(&ret).Error
        if err != nil {
            return err
        }
        return buildQueryFrom(tx, teacherId, models.Exam{}).Count(&num).Error
    })
    utils.PanicWhen(err)
    return
}

// GetExams searches the database for all exams.
// When any error occurs, it panics.
func GetExams() (ret []models.Exam) {
    utils.PanicWhen(db.Find(&ret).Error)
    return
}

// CreateExams stores given exam into the database,
// with their id property ignored and handled by the database. When any error occurs,
// it panics and none of the given exam will be created alone.
func CreateExams(exams []*models.Exam) {
    utils.PanicWhen(db.Create(&exams).Error)
}

// UpdateExams updates all the columns of all given exam,
// the record to be updated will be specified by given exam's id.
// If any id of given `exams` doesn't exist or has started (including ended), it refuses to proceed and throws an error.
// When any error occurs, it panics and none of the given exam will be updated alone.
func UpdateExams(exams []*models.Exam) {
    err := db.Transaction(func(tx *gorm.DB) error {
        tmpExam := &models.Exam{}
        for _, exam := range exams {
            // SELECT FOR UPDATE, make sure all the ids exist and not started
            err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
                Select("id").
                Where("id = ? AND CURTIME() < start_time", exam.ID).
                First(tmpExam).Error
            if err != nil {
                return err
            }
            err = tx.Where("id = ?", exam.ID).Updates(exam).Error
            if err != nil {
                return err
            }
        }
        return nil
    })
    utils.PanicWhen(err)
}

// DeleteExams deletes all the records whose id is in given `ids`.
// If any id in given `ids` doesn't exist or is active, it refuses to proceed and throws an error.
// When any error occurs, it panics and none of the given exam will be deleted alone.
func DeleteExams(ids []int) {
    err := db.Transaction(func(tx *gorm.DB) error {
        for _, id := range ids {
            // SELECT FOR UPDATE, make sure all the ids exist and not active
            err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
                Select("id").
                Where("id = ? AND (CURTIME() < start_time OR end_time < CURTIME())", id).
                First(&models.Exam{}).Error
            if err != nil {
                return err
            }
        }
        // batch delete
        return tx.Delete(&models.Exam{}, ids).Error
    })
    utils.PanicWhen(err)
}

func EnterExam(studentId string, examId int) {
    err := db.Transaction(func(tx *gorm.DB) error {
        var exam =models.Exam{}
        var err error

        // SELECT FOR UPDATE, make sure the exam exists and active
        err = tx.Clauses(clause.Locking{Strength: "SHARED"}).
            Select("id").
            Where("id = ? AND start_time <= CURTIME() AND CURTIME() < end_time", examId).
            First(&exam).Error
        if err != nil {
            return err
        }
        var session = models.ExamSession{
            ExamID:      examId,
            StudentID:   studentId,
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
                err = tx.Create(models.McqAnswer{
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
                err = tx.Create(models.MaqAnswer{
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
            err = tx.Select("id", "right_answer").Order("RAND()").Limit(int(exam.BfqNum)).Find(&bfqs).Error
            if err != nil {
                return err
            }
            for _, bfq := range bfqs {
                err = tx.Create(models.BfqAnswer{
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
            err = tx.Select("id", "right_answer").Order("RAND()").Limit(int(exam.TfqNum)).Find(&tfqs).Error
            if err != nil {
                return err
            }
            for _, tfq := range tfqs {
                err = tx.Create(models.TfqAnswer{
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
            err = tx.Select("id", "right_answer").Order("RAND()").Limit(int(exam.CrqNum)).Find(&crqs).Error
            if err != nil {
                return err
            }
            for _, crq := range crqs {
                err = tx.Create(models.CrqAnswer{
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
