package dao

import (
    "github.com/gonearewe/EasyTesting/models"
    "github.com/gonearewe/EasyTesting/utils"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

// GetMcqsBy searches the database for Multiple Choice Question (mcq) whose publisher teacher id (string) starts
// with `teacherId`, it only returns records in given `pageIndex` (1-based) in the increasing order of id.
// When any error occurs, it panics.
func GetMcqsBy(teacherId string, pageSize int, pageIndex int) (ret []*models.Mcq) {
    var err error
    if teacherId != "" {
        err = db.Limit(pageSize).Offset(pageSize*(pageIndex-1)).Find(&ret, "publisher_teacher_id LIKE ?",
            teacherId+"%").Error
    } else {
        err = db.Limit(pageSize).Offset(pageSize * (pageIndex - 1)).Find(&ret).Error
    }
    utils.PanicWhen(err)
    return
}

// CreateMcqs stores given Multiple Choice Question (mcq) into the database,
// with their id property ignored and handled by the database. When any error occurs,
// it panics and none of the given mcq will be created alone.
func CreateMcqs(questions []*models.Mcq) {
    utils.PanicWhen(db.Create(&questions).Error)
}

// UpdateMcqs updates all the columns of all given Multiple Choice Question (mcq),
// the record to be updated will be specified by given mcq's id.
// If any id of given `questions` doesn't exist, it refuses to proceed and throws an error.
// When any error occurs, it panics and none of the given mcq will be updated alone.
func UpdateMcqs(questions []*models.Mcq) {
    err := db.Transaction(func(tx *gorm.DB) error {
        tmpMcq := &models.Mcq{}
        for _, mcq := range questions {
            // SELECT FOR UPDATE, make sure all the ids exist
            err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
                Select("id").Where("id = ?", mcq.ID).First(tmpMcq).Error
            if err != nil {
                return err
            }
            err = tx.Where("id = ?", mcq.ID).Updates(mcq).Error
            if err != nil {
                return err
            }
        }
        return nil
    })
    utils.PanicWhen(err)
}

// DeleteMcqs deletes all the records whose id is in given `ids`.
// If any id in given `ids` doesn't exist, it refuses to proceed and throws an error.
// When any error occurs, it panics and none of the given mcq will be deleted alone.
func DeleteMcqs(ids []int) {
    err := db.Transaction(func(tx *gorm.DB) error {
        tmpMcq := &models.Mcq{}
        for _, id := range ids {
            // SELECT FOR UPDATE, make sure all the ids exist
            err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
                Select("id").Where("id = ?", id).First(tmpMcq).Error
            if err != nil {
                return err
            }
        }
        // batch delete
        return tx.Delete(&models.Mcq{}, ids).Error
    })
    utils.PanicWhen(err)
}
