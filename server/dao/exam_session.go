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
        var sessionId2ScoreMap = map[int]int{}
        for i, e := range examinees {
            ids[i] = e.ID
            sessionId2ScoreMap[e.ID] = 0
        }

        type overallStatistics struct {
            overallCorrectScore int
            overallScore int
        }
        var mcqId2StatisticsMap = map[int]*overallStatistics{}
        var maqId2StatisticsMap = map[int]*overallStatistics{}
        var bfqId2StatisticsMap = map[int]*overallStatistics{}
        var tfqId2StatisticsMap = map[int]*overallStatistics{}
        var crqId2StatisticsMap = map[int]*overallStatistics{}
        var cqId2StatisticsMap = map[int]*overallStatistics{}

        var mcqAns []*models.McqAnswer
        utils.PanicWhen(tx.Where("exam_session_id IN ?", ids).
            Clauses(clause.Locking{Strength: "SHARE"}).Find(&mcqAns).Error)
        for _,ans := range mcqAns{
            mcqId2StatisticsMap[ans.McqID] = &overallStatistics{}
        }

        var maqAns []*models.MaqAnswer
        utils.PanicWhen(tx.Where("exam_session_id IN ?", ids).
            Clauses(clause.Locking{Strength: "SHARE"}).Find(&maqAns).Error)
        for _,ans := range maqAns{
            maqId2StatisticsMap[ans.MaqID] = &overallStatistics{}
        }

        var bfqAns []*models.BfqAnswer
        utils.PanicWhen(tx.Where("exam_session_id IN ?", ids).
            Clauses(clause.Locking{Strength: "SHARE"}).Find(&bfqAns).Error)
        for _,ans := range bfqAns{
            bfqId2StatisticsMap[ans.BfqID] = &overallStatistics{}
        }

        var tfqAns []*models.TfqAnswer
        utils.PanicWhen(tx.Where("exam_session_id IN ?", ids).
            Clauses(clause.Locking{Strength: "SHARE"}).Find(&tfqAns).Error)
        for _,ans := range tfqAns{
            tfqId2StatisticsMap[ans.TfqID] = &overallStatistics{}
        }

        var crqAns []*models.CrqAnswer
        utils.PanicWhen(tx.Where("exam_session_id IN ?", ids).
            Clauses(clause.Locking{Strength: "SHARE"}).Find(&crqAns).Error)
        for _,ans := range crqAns {
            crqId2StatisticsMap[ans.CrqID] = &overallStatistics{}
        }

        var cqAns []*models.CqAnswer
        utils.PanicWhen(tx.Where("exam_session_id IN ?", ids).
            Clauses(clause.Locking{Strength: "SHARE"}).Find(&cqAns).Error)
        for _,ans := range cqAns{
            cqId2StatisticsMap[ans.CqID] = &overallStatistics{}
        }

        for _, e := range mcqAns {
            statistics := mcqId2StatisticsMap[e.McqID]
            statistics.overallScore += int(exam.McqScore) * 10
            if e.StudentAnswer == e.RightAnswer {
                // score stored in `exam` table is integer, but that in `exam_session` table is decimal,
                // so a factor of 10 is needed.
                sessionId2ScoreMap[e.ExamSessionID] += int(exam.McqScore) * 10
                statistics.overallCorrectScore += int(exam.McqScore) * 10
            }
        }
        for _, e := range maqAns {
            statistics := maqId2StatisticsMap[e.MaqID]
            statistics.overallScore += int(exam.MaqScore) * 10
            if utils.IsAnagram(e.StudentAnswer, e.RightAnswer) { // 全选对得满分
                sessionId2ScoreMap[e.ExamSessionID] += int(exam.MaqScore) * 10
                statistics.overallCorrectScore += int(exam.MaqScore) * 10
            }else if utils.Contains(e.RightAnswer,e.StudentAnswer) { // 漏选得半分
                sessionId2ScoreMap[e.ExamSessionID] += int(exam.MaqScore) * 10 / 2
                statistics.overallCorrectScore += int(exam.MaqScore) * 10 / 2
            }
        }
        for _, e := range bfqAns {
            var bfq models.Bfq
            utils.PanicWhen(tx.Select("blank_num", "answer_1", "answer_2", "answer_3").
                Where("id = ?", e.BfqID).First(&bfq).Error)
            statistics := bfqId2StatisticsMap[e.BfqID]
            statistics.overallScore += int(exam.BfqScore) * 10
            for i, ans := range [][]string{
                []string{e.StudentAnswer1, bfq.Answer1}, []string{e.StudentAnswer2, bfq.Answer2},
                []string{e.StudentAnswer3, bfq.Answer3},
            } {
                if i >= int(bfq.BlankNum) {
                    break
                }
                if ans[0] == ans[1] {
                    score := int(exam.BfqScore) * 10 / int(bfq.BlankNum)
                    sessionId2ScoreMap[e.ExamSessionID] += score
                    statistics.overallCorrectScore += score
                }
            }
        }
        for _, e := range tfqAns {
            statistics := tfqId2StatisticsMap[e.TfqID]
            statistics.overallScore += int(exam.TfqScore) * 10
            if e.StudentAnswer == e.RightAnswer {
                sessionId2ScoreMap[e.ExamSessionID] += int(exam.TfqScore) * 10
                statistics.overallCorrectScore += int(exam.TfqScore) * 10
            }
        }
        for _, e := range crqAns {
            var crq models.Crq
            utils.PanicWhen(tx.Select(
                "blank_num", "answer_1", "answer_2", "answer_3", "answer_4", "answer_5", "answer_6",
            ).Where("id = ?", e.CrqID).First(&crq).Error)
            statistics := crqId2StatisticsMap[e.CrqID]
            statistics.overallScore += int(exam.CrqScore) * 10
            for i, ans := range [][]string{
                []string{e.StudentAnswer1, crq.Answer1}, []string{e.StudentAnswer2, crq.Answer2},
                []string{e.StudentAnswer3, crq.Answer3}, []string{e.StudentAnswer4, crq.Answer4},
                []string{e.StudentAnswer5, crq.Answer5}, []string{e.StudentAnswer6, crq.Answer6},
            } {
                if i >= int(crq.BlankNum) {
                    break
                }
                if ans[0] == ans[1] {
                    score := int(exam.CrqScore) * 10 / int(crq.BlankNum)
                    sessionId2ScoreMap[e.ExamSessionID] += score
                    statistics.overallCorrectScore += score
                }
            }
        }
        for _, e := range cqAns {
            statistics := cqId2StatisticsMap[e.CqID]
            statistics.overallScore += int(exam.CqScore) * 10
            if e.IsAnswerRight {
                sessionId2ScoreMap[e.ExamSessionID] += int(exam.CqScore) * 10
                statistics.overallCorrectScore += int(exam.CqScore) * 10
            }
        }

        for id, score := range sessionId2ScoreMap {
            utils.PanicWhen(tx.Model(&models.ExamSession{}).Where("id = ?", id).Update("score", score).Error)
        }
        tmp:= []string{"mcq","maq","bfq","tfq","crq","cq"}
        for i, m := range []map[int]*overallStatistics{mcqId2StatisticsMap,maqId2StatisticsMap,
            bfqId2StatisticsMap,tfqId2StatisticsMap,crqId2StatisticsMap,cqId2StatisticsMap} {
            for id, statistics := range m {
                utils.PanicWhen(tx.Exec(
                    "UPDATE "+tmp[i]+" SET overall_score = overall_score + ?, "+
                        "overall_correct_score = overall_correct_score + ? WHERE id = ?",
                    statistics.overallScore, statistics.overallCorrectScore, id).Error)
            }
        }
        utils.PanicWhen(tx.Model(&models.Exam{}).Where("id = ?", examId).Update("scores_calculated", true).Error)
        return nil
    })
    utils.PanicWhen(err)
}
