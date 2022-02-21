package dao

import (
    "github.com/gonearewe/EasyTesting/models"
    "github.com/gonearewe/EasyTesting/utils"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

// GetCrqsBy searches the database for Code Reading Question (crq) whose publisher teacher id (string) starts
// with `teacherId`, it only returns records in given `pageIndex` (1-based) in the increasing order of id.
// When any error occurs, it panics.
func GetCrqsBy(teacherId string, pageSize int, pageIndex int) (ret []*models.Crq, num int64) {
    err := db.Transaction(func(tx *gorm.DB) error {
        err := buildCrqQueryFrom(tx, teacherId).
            Limit(pageSize).Offset(pageSize * (pageIndex - 1)).
            Find(&ret).Error
        if err != nil {
            return err
        }
        return buildCrqQueryFrom(tx, teacherId).Count(&num).Error
    })
    utils.PanicWhen(err)
    return
}

func buildCrqQueryFrom(tx *gorm.DB, teacherId string) *gorm.DB {
    var filtered = tx.Model(&models.Crq{})
    if teacherId != "" {
        filtered = filtered.Where("publisher_teacher_id LIKE ?", teacherId+"%")
    }
    return filtered
}

// CreateCrqs stores given Code Reading Question (crq) into the database,
// with their id property ignored and handled by the database. When any error occurs,
// it panics and none of the given crq will be created alone.
func CreateCrqs(questions []*models.Crq) {
    utils.PanicWhen(db.Create(&questions).Error)
}

// UpdateCrqById updates all the columns of given Code Reading Question (crq),
// the record to be updated will be specified by given crq's id.
// When any error occurs, it panics and the given crq will not be updated.
func UpdateCrqById(question *models.Crq) {
    err := db.Model(question).Where("id = ?", question.ID).Updates(
        map[string]interface{}{
            "stem":question.Stem,
            "blank_num":question.BlankNum,
            "answer_1":question.Answer1,
            "answer_2":question.Answer2,
            "answer_3":question.Answer3,
            "answer_4":question.Answer4,
            "answer_5":question.Answer5,
            "answer_6":question.Answer6,
            "overall_score":       0,
            "overall_correct_score":      0,
        }).Error
    utils.PanicWhen(err)
}

// DeleteCrqs deletes all the records whose id is in given `ids`.
// If any id in given `ids` doesn't exist, it refuses to proceed and throws an error.
// When any error occurs, it panics and none of the given crq will be deleted alone.
func DeleteCrqs(ids []int) {
    err := db.Transaction(func(tx *gorm.DB) error {
        tmpCrq := &models.Crq{}
        for _, id := range ids {
            // SELECT FOR UPDATE, make sure all the ids exist
            err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
                Select("id").Where("id = ?", id).First(tmpCrq).Error
            if err != nil {
                return err
            }
        }
        // batch delete
        return tx.Delete(&models.Crq{}, ids).Error
    })
    utils.PanicWhen(err)
}
