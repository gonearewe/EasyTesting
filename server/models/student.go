package models

import (
	"time"
)

// Student [...]
type Student struct {
	ID            int       `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                                           // 用作主键
	StudentID     string    `gorm:"unique;index:student_id_2;column:student_id;type:varchar(10);not null" json:"student_id"`        // 学号
	Name          string    `gorm:"index:student_id_2;index:class_id;index:name;column:name;type:varchar(50);not null" json:"name"` // 姓名
	ClassID       string    `gorm:"index:student_id_2;index:class_id;column:class_id;type:varchar(10);not null" json:"class_id"`    // 班号
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	LastUpdatedAt time.Time `gorm:"column:last_updated_at;type:datetime" json:"last_updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *Student) TableName() string {
	return "student"
}
