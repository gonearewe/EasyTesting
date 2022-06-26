package dao

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/gonearewe/EasyTesting/models"
	"github.com/gonearewe/EasyTesting/utils"
)

// GetCqsBy searches the database for Coding Question (cq) whose publisher teacher id (string) starts
// with `teacherId`, it only returns records in given `pageIndex` (1-based) in the increasing order of id.
// When any error occurs, it panics.
func GetCqsBy(teacherId string, pageSize int, pageIndex int) (ret []*models.Cq, num int64) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := buildQueryFrom(tx, teacherId, &models.Cq{}).
			Limit(pageSize).Offset(pageSize * (pageIndex - 1)).
			Find(&ret).Error
		if err != nil {
			return err
		}
		return buildQueryFrom(tx, teacherId, &models.Cq{}).Count(&num).Error
	})
	utils.PanicWhen(err)
	return
}

// CreateCqs stores given Coding Question (cq) into the database,
// with their id property ignored and handled by the database. When any error occurs,
// it panics and none of the given cq will be created alone.
func CreateCqs(questions []*models.Cq) {
	utils.PanicWhen(db.Create(&questions).Error)
}

// UpdateCqById updates all the columns of given Coding Question (cq),
// the record to be updated will be specified by given cq's id.
// When any error occurs, it panics and the given cq will not be updated.
func UpdateCqById(question *models.Cq) {
	err := db.Model(question).Where("id = ?", question.ID).Updates(
		map[string]interface{}{
			"stem":                  question.Stem,
			"input":                 question.Input,
			"output":                question.Output,
			"is_input_from_file":    question.IsInputFromFile,
			"is_output_to_file":     question.IsOutputToFile,
			"template":              question.Template,
			"overall_score":         0,
			"overall_correct_score": 0,
		}).Error
	utils.PanicWhen(err)
}

// DeleteCqs deletes all the records whose id is in given `ids`.
// If any id in given `ids` doesn't exist, it refuses to proceed and throws an error.
// When any error occurs, it panics and none of the given cq will be deleted alone.
func DeleteCqs(ids []int) {
	err := db.Transaction(func(tx *gorm.DB) error {
		tmpCq := &models.Cq{}
		for _, id := range ids {
			// SELECT FOR UPDATE, make sure all the ids exist
			err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				Select("id").Where("id = ?", id).First(tmpCq).Error
			if err != nil {
				return err
			}
		}
		// batch delete
		return tx.Delete(&models.Cq{}, ids).Error
	})
	utils.PanicWhen(err)
}
