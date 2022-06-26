package dao

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/gonearewe/EasyTesting/models"
	"github.com/gonearewe/EasyTesting/utils"
)

// GetTfqsBy searches the database for True False Question (tfq) whose publisher teacher id (string) starts
// with `teacherId`, it only returns records in given `pageIndex` (1-based) in the increasing order of id,
// but plus the total number of all filtered.
// When any error occurs, it panics.
func GetTfqsBy(teacherId string, pageSize int, pageIndex int) (ret []*models.Tfq, num int64) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := buildQueryFrom(tx, teacherId, &models.Tfq{}).
			Limit(pageSize).Offset(pageSize * (pageIndex - 1)).
			Find(&ret).Error
		if err != nil {
			return err
		}
		return buildQueryFrom(tx, teacherId, &models.Tfq{}).Count(&num).Error
	})
	utils.PanicWhen(err)
	return
}

// CreateTfqs stores given True False Question (tfq) into the database,
// with their id property ignored and handled by the database. When any error occurs,
// it panics and none of the given tfq will be created alone.
func CreateTfqs(questions []*models.Tfq) {
	utils.PanicWhen(db.Create(&questions).Error)
}

// UpdateTfqById updates all the columns of given True False Question (tfq),
// the record to be updated will be specified by given tfq's id.
// When any error occurs, it panics and the given tfq will not be updated.
func UpdateTfqById(question *models.Tfq) {
	err := db.Model(question).Where("id = ?", question.ID).Updates(
		map[string]interface{}{
			"stem":                  question.Stem,
			"answer":                question.Answer,
			"overall_score":         0,
			"overall_correct_score": 0,
		}).Error
	utils.PanicWhen(err)
}

// DeleteTfqs deletes all the records whose id is in given `ids`.
// If any id in given `ids` doesn't exist, it refuses to proceed and throws an error.
// When any error occurs, it panics and none of the given tfq will be deleted alone.
func DeleteTfqs(ids []int) {
	err := db.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			// SELECT FOR UPDATE, make sure all the ids exist
			err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				Select("id").Where("id = ?", id).First(&models.Tfq{}).Error
			if err != nil {
				return err
			}
		}
		// batch delete
		return tx.Delete(&models.Tfq{}, ids).Error
	})
	utils.PanicWhen(err)
}
