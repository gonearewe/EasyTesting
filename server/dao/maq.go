package dao

import (
    "github.com/gonearewe/EasyTesting/models"
    "github.com/gonearewe/EasyTesting/utils"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

// GetMaqsBy searches the database for Multiple Answer Question (maq) whose publisher teacher id (string) starts
// with `teacherId`, it only returns records in given `pageIndex` (1-based) in the increasing order of id.
// When any error occurs, it panics.
func GetMaqsBy(teacherId string, pageSize int, pageIndex int) (ret []*models.Maq, num int64) {
    err := db.Transaction(func(tx *gorm.DB) error {
        err := buildMaqQueryFrom(tx, teacherId).
            Limit(pageSize).Offset(pageSize * (pageIndex - 1)).
            Find(&ret).Error
        if err != nil {
            return err
        }
        return buildMaqQueryFrom(tx, teacherId).Count(&num).Error
    })
    utils.PanicWhen(err)
    return
}

func buildMaqQueryFrom(tx *gorm.DB, teacherId string) *gorm.DB {
    var filtered = tx.Model(&models.Maq{})
    if teacherId != "" {
        filtered = filtered.Where("publisher_teacher_id LIKE ?", teacherId+"%")
    }
    return filtered
}

// CreateMaqs stores given Multiple Answer Question (maq) into the database,
// with their id property ignored and handled by the database. When any error occurs,
// it panics and none of the given maq will be created alone.
func CreateMaqs(questions []*models.Maq) {
    utils.PanicWhen(db.Create(&questions).Error)
}

// UpdateMaqById updates all the columns of given Multiple Answer Question (maq),
// the record to be updated will be specified by given maq's id.
// When any error occurs, it panics and the given maq will not be updated.
func UpdateMaqById(question *models.Maq) {
    err := db.Model(question).Where("id = ?", question.ID).Updates(
        map[string]interface{}{
            "stem":question.Stem,
            "choice_1":question.Choice1,
            "choice_2":question.Choice2,
            "choice_3":question.Choice3,
            "choice_4":question.Choice4,
            "choice_5":question.Choice5,
            "choice_6":question.Choice6,
            "choice_7":question.Choice7,
            "right_answer":question.RightAnswer,
            "overall_score":       0,
            "overall_correct_score":      0,
        }).Error
    utils.PanicWhen(err)
}

// DeleteMaqs deletes all the records whose id is in given `ids`.
// If any id in given `ids` doesn't exist, it refuses to proceed and throws an error.
// When any error occurs, it panics and none of the given maq will be deleted alone.
func DeleteMaqs(ids []int) {
    err := db.Transaction(func(tx *gorm.DB) error {
        tmpMaq := &models.Maq{}
        for _, id := range ids {
            // SELECT FOR UPDATE, make sure all the ids exist
            err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
                Select("id").Where("id = ?", id).First(tmpMaq).Error
            if err != nil {
                return err
            }
        }
        // batch delete
        return tx.Delete(&models.Maq{}, ids).Error
    })
    utils.PanicWhen(err)
}
