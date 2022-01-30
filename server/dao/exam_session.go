package dao

import (
    "errors"
    "strconv"
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
    err = db.Select("id", "student_name", "start_time", "time_allowed").
        Where("exam_id = ? AND student_id = ?", examId, studentId).First(&ret).Error
    return
}

func GetEndedExamSessionsBy(examId int) (ret []*models.ExamSession) {
    // `answer_sheet` field is excluded.
    db.Select("student_id", "student_name", "start_time", "end_time", "score").
        Where("exam_id = ?", examId).Find(&ret)
    return
}

func EnterExam(studentId string, studentName string, examId int) {
    err := db.Transaction(func(tx *gorm.DB) error {
        var exam = models.Exam{}
        var err error

        // SELECT FOR SHARE, make sure the exam exists and active
        err = tx.Clauses(clause.Locking{Strength: "SHARE"}).
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
            StudentName: studentName,
            TimeAllowed: exam.TimeAllowed,
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
                    CqID:          cq.ID,
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
            var cqAnswers []*models.CqAnswer
            err = tx.Select("cq_id").Where("exam_session_id = ?", examSessionId).Find(&cqAnswers).Error
            if err != nil {
                return err
            }
            var cqs []*models.Cq
            var cqIds = make([]int, len(cqAnswers))
            for i, a := range cqAnswers {
                cqIds[i] = a.CqID
            }
            err = tx.Where("id IN ?", cqIds).Find(&cqs).Error
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

func SubmitMyAnswers(examSessionId int, mcqs []*models.McqAnswer, maqs []*models.MaqAnswer, bfqs []*models.BfqAnswer,
    tfqs []*models.TfqAnswer, crqs []*models.CrqAnswer, cqs []*models.CqAnswer) {
    // it's not necessary to execute UPDATE operations in a transaction,
    // since there's no data consistency constrain between different type of answers

    errOccurred := false

    for _, mcq := range mcqs {
        err := db.Model(&models.McqAnswer{}).Where("mcq_id = ?", mcq.McqID).
            Where("exam_session_id = ?", examSessionId).Update("student_answer", mcq.StudentAnswer).Error
        errOccurred = errOccurred || err != nil
    }

    for _, maq := range maqs {
        err := db.Model(&models.MaqAnswer{}).Where("maq_id = ?", maq.MaqID).
            Where("exam_session_id = ?", examSessionId).Update("student_answer", maq.StudentAnswer).Error
        errOccurred = errOccurred || err != nil
    }

    for _, bfq := range bfqs {
        err := db.Model(&models.BfqAnswer{}).Where("bfq_id = ?", bfq.BfqID).
            Where("exam_session_id = ?", examSessionId).Updates(
            map[string]interface{}{
                "student_answer_1": bfq.StudentAnswer1,
                "student_answer_2": bfq.StudentAnswer2,
                "student_answer_3": bfq.StudentAnswer3,
            },
        ).Error
        errOccurred = errOccurred || err != nil
    }

    for _, tfq := range tfqs {
        err := db.Model(&models.TfqAnswer{}).Where("tfq_id = ?", tfq.TfqID).
            Where("exam_session_id = ?", examSessionId).Update("student_answer", tfq.StudentAnswer).Error
        errOccurred = errOccurred || err != nil
    }

    for _, crq := range crqs {
        err := db.Model(&models.CrqAnswer{}).Where("crq_id = ?", crq.CrqID).
            Where("exam_session_id = ?", examSessionId).Updates(
            map[string]interface{}{
                "student_answer_1": crq.StudentAnswer1,
                "student_answer_2": crq.StudentAnswer2,
                "student_answer_3": crq.StudentAnswer3,
                "student_answer_4": crq.StudentAnswer4,
                "student_answer_5": crq.StudentAnswer5,
                "student_answer_6": crq.StudentAnswer6,
            },
        ).Error
        errOccurred = errOccurred || err != nil
    }

    for _, cq := range cqs {
        err := db.Model(&models.CqAnswer{}).Where("cq_id = ?", cq.CqID).
            Where("exam_session_id = ?", examSessionId).Updates(
            map[string]interface{}{
                "student_answer":  cq.StudentAnswer,
                "is_answer_right": cq.IsAnswerRight,
            }).Error
        errOccurred = errOccurred || err != nil
    }

    // record answer submitting time
    err := db.Model(&models.ExamSession{}).Where("id = ?", examSessionId).Update("end_time", time.Now()).Error
    errOccurred = errOccurred || err != nil
    // TODO: log errors
    if errOccurred {
        panic(errors.New("some error occur when submitting answers"))
    }
}

func CalculateScores(examId int) {
    err := db.Transaction(func(tx *gorm.DB) (err error) {
        defer func() {
            if e := recover(); e != nil {
                err = e.(error)
            }
        }()

        var exam models.Exam
        utils.PanicWhen(tx.Where("id = ?", examId).Clauses(clause.Locking{Strength: "SHARE"}).First(&exam).Error)
        if exam.EndTime.After(time.Now()) || exam.ScoresCalculated {
            return errors.New("scores of this exam have already been calculated or can only be calculated later: " +
                strconv.Itoa(examId))
        }

        var examinees []models.ExamSession
        utils.PanicWhen(tx.Select("id").Where("exam_id = ?", examId).
            Clauses(clause.Locking{Strength: "SHARE"}).Find(&examinees).Error)
        ids := make([]int, len(examinees))
        var idScoreMap = map[int]int{}
        for i, e := range examinees {
            ids[i] = e.ID
            idScoreMap[e.ID] = 0
        }

        var mcqAns []*models.McqAnswer
        utils.PanicWhen(tx.Select("exam_session_id", "right_answer", "student_answer").Where("exam_session_id IN ?", ids).
            Clauses(clause.Locking{Strength: "SHARE"}).Find(&mcqAns).Error)

        var maqAns []*models.MaqAnswer
        utils.PanicWhen(tx.Select("exam_session_id", "right_answer", "student_answer").Where("exam_session_id IN ?", ids).
            Clauses(clause.Locking{Strength: "SHARE"}).Find(&maqAns).Error)

        var bfqAns []*models.BfqAnswer
        utils.PanicWhen(tx.Where("exam_session_id IN ?", ids).
            Clauses(clause.Locking{Strength: "SHARE"}).Find(&bfqAns).Error)

        var tfqAns []*models.TfqAnswer
        utils.PanicWhen(tx.Select("exam_session_id", "right_answer", "student_answer").Where("exam_session_id IN ?", ids).
            Clauses(clause.Locking{Strength: "SHARE"}).Find(&tfqAns).Error)

        var crqAns []*models.CrqAnswer
        utils.PanicWhen(tx.Where("exam_session_id IN ?", ids).
            Clauses(clause.Locking{Strength: "SHARE"}).Find(&crqAns).Error)

        var cqAns []*models.CqAnswer
        utils.PanicWhen(tx.Select("exam_session_id", "is_answer_right").Where("exam_session_id IN ?", ids).
            Clauses(clause.Locking{Strength: "SHARE"}).Find(&cqAns).Error)

        for _, e := range mcqAns {
            if e.StudentAnswer == e.RightAnswer {
                // score stored in `exam` table is integer, but that in `exam_session` table is decimal,
                // so a factor of 10 is needed.
                idScoreMap[e.ExamSessionID] += int(exam.McqScore) * 10
            }
        }
        for _, e := range maqAns {
            if utils.IsAnagram(e.StudentAnswer, e.RightAnswer) {
                idScoreMap[e.ExamSessionID] += int(exam.MaqScore) * 10
            }
        }
        for _, e := range bfqAns {
            var bfq models.Bfq
            utils.PanicWhen(tx.Select("blank_num", "answer_1", "answer_2", "answer_3").
                Where("id = ?", e.BfqID).First(&bfq).Error)
            for i, ans := range [][]string{
                []string{e.StudentAnswer1, bfq.Answer1}, []string{e.StudentAnswer2, bfq.Answer2},
                []string{e.StudentAnswer3, bfq.Answer3},
            } {
                if i >= int(bfq.BlankNum) {
                    break
                }
                if ans[0] == ans[1] {
                    idScoreMap[e.ExamSessionID] += int(exam.BfqScore) * 10 / int(bfq.BlankNum)
                }
            }
        }
        for _, e := range tfqAns {
            if e.StudentAnswer == e.RightAnswer {
                idScoreMap[e.ExamSessionID] += int(exam.TfqScore) * 10
            }
        }
        for _, e := range crqAns {
            var crq models.Crq
            utils.PanicWhen(tx.Select(
                "blank_num", "answer_1", "answer_2", "answer_3", "answer_4", "answer_5", "answer_6",
            ).Where("id = ?", e.CrqID).First(&crq).Error)
            for i, ans := range [][]string{
                []string{e.StudentAnswer1, crq.Answer1}, []string{e.StudentAnswer2, crq.Answer2},
                []string{e.StudentAnswer3, crq.Answer3}, []string{e.StudentAnswer4, crq.Answer4},
                []string{e.StudentAnswer5, crq.Answer5}, []string{e.StudentAnswer6, crq.Answer6},
            } {
                if i >= int(crq.BlankNum) {
                    break
                }
                if ans[0] == ans[1] {
                    idScoreMap[e.ExamSessionID] += int(exam.CrqScore) * 10 / int(crq.BlankNum)
                }
            }
        }
        for _, e := range cqAns {
            if e.IsAnswerRight {
                idScoreMap[e.ExamSessionID] += int(exam.CqScore) * 10
            }
        }

        for id, score := range idScoreMap {
            utils.PanicWhen(tx.Model(&models.ExamSession{}).Where("id = ?", id).Update("score", score).Error)
        }
        utils.PanicWhen(tx.Model(&models.Exam{}).Where("id = ?", examId).Update("scores_calculated", true).Error)
        return nil
    })
    utils.PanicWhen(err)
}
