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

// GetMcqNumBy searches the database for the number of Multiple Choice Question (mcq) whose publisher teacher id (
// string) starts with `teacherId`.
// When any error occurs, it panics.
func GetMcqNumBy(teacherId string) (num int64) {
    var err error
    if teacherId != "" {
        err = db.Model(&models.Mcq{}).Where("publisher_teacher_id LIKE ?", teacherId+"%").Count(&num).Error
    } else {
        err = db.Model(&models.Mcq{}).Count(&num).Error
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

// UpdateMcqById updates all the columns of given Multiple Choice Question (mcq),
// the record to be updated will be specified by given mcq's id.
// When any error occurs, it panics and the given mcq will not be updated.
func UpdateMcqById(question *models.Mcq) {
    err := db.Where("id = ?", question.ID).Updates(question).Error
    utils.PanicWhen(err)
}

// DeleteMcqs deletes all the records whose id is in given `ids`.
// If any id in given `ids` doesn't exist, it refuses to proceed and throws an error.
// When any error occurs, it panics and none of the given mcq will be deleted alone.
func DeleteMcqs(ids []int) {
    err := db.Transaction(func(tx *gorm.DB) error {
        for _, id := range ids {
            // SELECT FOR UPDATE, make sure all the ids exist
            err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
                Select("id").Where("id = ?", id).First(&models.Mcq{}).Error
            if err != nil {
                return err
            }
        }
        // batch delete
        return tx.Delete(&models.Mcq{}, ids).Error
    })
    utils.PanicWhen(err)
}
