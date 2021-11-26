package models

// McqAnswer [...]
type McqAnswer struct {
	ID            int    `gorm:"primaryKey;column:id"`   // 用作主键
	McqID         int    `gorm:"column:mcq_id"`          // 连接 mcq
	ExamSessionID int    `gorm:"column:exam_session_id"` // 连接 exam_session
	RightAnswer   string `gorm:"column:right_answer"`    // 正确答案，与 mcq 中同名字段保持一致
	StudentAnswer string `gorm:"column:student_answer"`  // 学生的答案
}

// TableName get sql table name.获取数据库表名
func (m *McqAnswer) TableName() string {
	return "mcq_answer"
}
