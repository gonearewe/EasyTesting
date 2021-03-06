package dao

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/gonearewe/EasyTesting/models"
	"github.com/gonearewe/EasyTesting/utils"
)

// GetMcqsBy searches the database for Multiple Choice Question (mcq) whose publisher teacher id (string) starts
// with `teacherId`, it only returns records in given `pageIndex` (1-based) in the increasing order of id,
// but plus the total number of all filtered.
// When any error occurs, it panics.
func GetMcqsBy(teacherId string, pageSize int, pageIndex int) (ret []*models.Mcq, num int64) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := buildQueryFrom(tx, teacherId, &models.Mcq{}).
			Limit(pageSize).Offset(pageSize * (pageIndex - 1)).
			Find(&ret).Error
		if err != nil {
			return err
		}
		return buildQueryFrom(tx, teacherId, &models.Mcq{}).Count(&num).Error
	})
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
	err := db.Model(question).Where("id = ?", question.ID).Updates(
		map[string]interface{}{
			"stem":                  question.Stem,
			"choice_1":              question.Choice1,
			"choice_2":              question.Choice2,
			"choice_3":              question.Choice3,
			"choice_4":              question.Choice4,
			"right_answer":          question.RightAnswer,
			"overall_score":         0,
			"overall_correct_score": 0,
		}).Error
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
