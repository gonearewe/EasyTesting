package dao

import (
    "github.com/gonearewe/EasyTesting/utils"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

func deleteBy(ids []int, tmpModel interface{})  {
    err := db.Transaction(func(tx *gorm.DB) error {
        for _, id := range ids {
            // SELECT FOR UPDATE, make sure all the ids exist
            err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
                Select("id").Where("id = ?", id).First(&tmpModel).Error
            if err != nil {
                return err
            }
        }
        // batch delete
        return tx.Delete(&tmpModel, ids).Error
    })
    utils.PanicWhen(err)
}

func buildQueryFrom(tx *gorm.DB, publisherTeacherId string, tmpModel interface{}) *gorm.DB {
    var filtered = tx.Model(&tmpModel)
    if publisherTeacherId != "" {
        filtered = filtered.Where("publisher_teacher_id LIKE ?", publisherTeacherId+"%")
    }
    return filtered
}
