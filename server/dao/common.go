package dao

import (
    "gorm.io/gorm"
)

func buildQueryFrom(tx *gorm.DB, publisherTeacherId string, tmpModel interface{}) *gorm.DB {
    var filtered = tx.Model(&tmpModel)
    if publisherTeacherId != "" {
        filtered = filtered.Where("publisher_teacher_id LIKE ?", publisherTeacherId+"%")
    }
    return filtered
}
