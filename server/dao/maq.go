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
func GetMaqsBy(teacherId string, pageSize int, pageIndex int) (ret []*models.Maq) {
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
    err := db.Where("id = ?", question.ID).Updates(question).Error
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
