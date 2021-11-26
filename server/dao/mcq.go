package dao

import (
	"github.com/gonearewe/EasyTesting/models"
	"github.com/gonearewe/EasyTesting/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetMcqsBy(teacherId string, pageSize int, pageIndex int) (ret []*models.Mcq) {
	var err error
	if teacherId != "" {
		err = db.Limit(pageSize).Offset(pageSize*(pageIndex-1)).Find(&ret, "teacher_id LIKE ?", teacherId+"%").Error
	} else {
		err = db.Limit(pageSize).Offset(pageSize * (pageIndex - 1)).Find(&ret).Error
	}
	utils.PanicWhen(err)
	return
}

func CreateMcqs(questions []*models.Mcq) {
	utils.PanicWhen(db.Create(&questions).Error)
}

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
