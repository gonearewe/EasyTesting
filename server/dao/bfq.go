package dao

import (
    "github.com/gonearewe/EasyTesting/models"
    "github.com/gonearewe/EasyTesting/utils"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

// GetBfqsBy searches the database for Blank Filling Question (bfq) whose publisher teacher id (string) starts
// with `teacherId`, it only returns records in given `pageIndex` (1-based) in the increasing order of id.
// When any error occurs, it panics.
func GetBfqsBy(teacherId string, pageSize int, pageIndex int) (ret []*models.Bfq, num int64) {
    err := db.Transaction(func(tx *gorm.DB) error {
        err := buildBfqQueryFrom(tx, teacherId).
            Limit(pageSize).Offset(pageSize * (pageIndex - 1)).
            Find(&ret).Error
        if err != nil {
            return err
        }
        return buildBfqQueryFrom(tx, teacherId).Count(&num).Error
    })
    utils.PanicWhen(err)
    return
}

func buildBfqQueryFrom(tx *gorm.DB, teacherId string) *gorm.DB {
    var filtered = tx.Model(&models.Bfq{})
    if teacherId != "" {
        filtered = filtered.Where("publisher_teacher_id LIKE ?", teacherId+"%")
    }
    return filtered.Order("id")
}

// CreateBfqs stores given Blank Filling Question (bfq) into the database,
// with their id property ignored and handled by the database. When any error occurs,
// it panics and none of the given bfq will be created alone.
func CreateBfqs(questions []*models.Bfq) {
    utils.PanicWhen(db.Create(&questions).Error)
}

// UpdateBfqById updates all the columns of given Blank Filling Question (bfq),
// the record to be updated will be specified by given bfq's id.
// When any error occurs, it panics and the given bfq will not be updated.
func UpdateBfqById(question *models.Bfq) {
    err := db.Model(question).Where("id = ?", question.ID).Updates(
        map[string]interface{}{
            "stem":question.Stem,
            "blank_num":question.BlankNum,
            "answer_1":question.Answer1,
            "answer_2":question.Answer2,
            "answer_3":question.Answer3,
            "overall_score":       0,
            "overall_correct_score":      0,
        }).Error
    utils.PanicWhen(err)
}

// DeleteBfqs deletes all the records whose id is in given `ids`.
// If any id in given `ids` doesn't exist, it refuses to proceed and throws an error.
// When any error occurs, it panics and none of the given bfq will be deleted alone.
func DeleteBfqs(ids []int) {
    err := db.Transaction(func(tx *gorm.DB) error {
        tmpBfq := &models.Bfq{}
        for _, id := range ids {
            // SELECT FOR UPDATE, make sure all the ids exist
            err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
                Select("id").Where("id = ?", id).First(tmpBfq).Error
            if err != nil {
                return err
            }
        }
        // batch delete
        return tx.Delete(&models.Bfq{}, ids).Error
    })
    utils.PanicWhen(err)
}
