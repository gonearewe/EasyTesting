package dao

import (
    "github.com/gonearewe/EasyTesting/models"
    "github.com/gonearewe/EasyTesting/utils"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

// GetExamsBy searches the database for exam whose publisher teacher id (string) starts
// with `teacherId`, it only returns records in given `pageIndex` (1-based) in the increasing order of id,
// but plus the total number of all filtered.
// When any error occurs, it panics.
func GetExamsBy(teacherId string, pageSize int, pageIndex int) (ret []*models.Exam,num int64) {
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
// If any id of given `exams` doesn't exist, it refuses to proceed and throws an error.
// When any error occurs, it panics and none of the given exam will be updated alone.
func UpdateExams(exams []*models.Exam) {
    err := db.Transaction(func(tx *gorm.DB) error {
        tmpExam := &models.Exam{}
        for _, exam := range exams {
            // SELECT FOR UPDATE, make sure all the ids exist
            err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
                Select("id").Where("id = ?", exam.ID).First(tmpExam).Error
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
// If any id in given `ids` doesn't exist, it refuses to proceed and throws an error.
// When any error occurs, it panics and none of the given exam will be deleted alone.
func DeleteExams(ids []int) {
    deleteBy(ids, models.Exam{})
}
