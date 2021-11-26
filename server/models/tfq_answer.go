package models

// TfqAnswer [...]
type TfqAnswer struct {
	ID            int  `gorm:"primaryKey;column:id"`   // 用作主键
	TfqID         int  `gorm:"column:tfq_id"`          // 连接 tfq
	ExamSessionID int  `gorm:"column:exam_session_id"` // 连接 exam_session
	RightAnswer   bool `gorm:"column:right_answer"`    // 正确答案，与 tfq 中同名字段保持一致
	StudentAnswer bool `gorm:"column:student_answer"`  // 学生的答案
}

// TableName get sql table name.获取数据库表名
func (m *TfqAnswer) TableName() string {
	return "tfq_answer"
}
