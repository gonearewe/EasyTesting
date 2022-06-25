package models

import (
	"time"
)

// CqAnswer [...]
type CqAnswer struct {
	ID            int       `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                                                        // 用作主键
	CqID          int       `gorm:"uniqueIndex:cq_id;column:cq_id;type:int(10);not null" json:"cq_id"`                                           // 连接 cq
	ExamSessionID int       `gorm:"uniqueIndex:cq_id;index:exam_session_id;column:exam_session_id;type:int(10);not null" json:"exam_session_id"` // 连接 exam_session
	StudentAnswer string    `gorm:"column:student_answer;type:text" json:"student_answer"`                                                       // 学生的答案，即代码
	RightOutput   string    `gorm:"column:right_output;type:text;not null" json:"right_output"`                                                  // 程序的正确输出，与 cq 中 output 字段保持一致
	StudentOutput string    `gorm:"column:student_output;type:text" json:"student_output"`                                                       // 学生代码的实际输出
	IsAnswerRight bool      `gorm:"column:is_answer_right;type:tinyint(1);default:0" json:"is_answer_right"`                                     // 学生的代码是否正确，即 right_output 是否等于 student_output，由触发器计算
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	LastUpdatedAt time.Time `gorm:"column:last_updated_at;type:datetime" json:"last_updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *CqAnswer) TableName() string {
	return "cq_answer"
}
