package models

import (
	"time"
)

// McqAnswer [...]
type McqAnswer struct {
	ID            int       `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                                                         // 用作主键
	McqID         int       `gorm:"uniqueIndex:mcq_id;column:mcq_id;type:int(10);not null" json:"mcq_id"`                                         // 连接 mcq
	ExamSessionID int       `gorm:"uniqueIndex:mcq_id;index:exam_session_id;column:exam_session_id;type:int(10);not null" json:"exam_session_id"` // 连接 exam_session
	RightAnswer   string    `gorm:"column:right_answer;type:char(1);not null" json:"right_answer"`                                                // 正确答案，与 mcq 中同名字段保持一致
	StudentAnswer string    `gorm:"column:student_answer;type:char(1)" json:"student_answer"`                                                     // 学生的答案
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	LastUpdatedAt time.Time `gorm:"column:last_updated_at;type:datetime" json:"last_updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *McqAnswer) TableName() string {
	return "mcq_answer"
}
