package models

import (
	"time"
)

// TfqAnswer [...]
type TfqAnswer struct {
	ID            int       `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                                                         // 用作主键
	TfqID         int       `gorm:"uniqueIndex:tfq_id;column:tfq_id;type:int(10);not null" json:"tfq_id"`                                         // 连接 tfq
	ExamSessionID int       `gorm:"uniqueIndex:tfq_id;index:exam_session_id;column:exam_session_id;type:int(10);not null" json:"exam_session_id"` // 连接 exam_session
	RightAnswer   bool      `gorm:"column:right_answer;type:tinyint(1);not null" json:"right_answer"`                                             // 正确答案，与 tfq 中同名字段保持一致
	StudentAnswer bool      `gorm:"column:student_answer;type:tinyint(1)" json:"student_answer"`                                                  // 学生的答案
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	LastUpdatedAt time.Time `gorm:"column:last_updated_at;type:datetime" json:"last_updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *TfqAnswer) TableName() string {
	return "tfq_answer"
}
