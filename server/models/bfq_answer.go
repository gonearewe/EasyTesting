package models

// BfqAnswer [...]
type BfqAnswer struct {
	ID             int    `gorm:"primaryKey;column:id"`    // 用作主键
	BfqID          int    `gorm:"column:bfq_id"`           // 连接 bfq
	ExamSessionID  int    `gorm:"column:exam_session_id"`  // 连接 exam_session
	StudentAnswer1 string `gorm:"column:student_answer_1"` // 学生的答案
	StudentAnswer2 string `gorm:"column:student_answer_2"` // 学生的答案
	StudentAnswer3 string `gorm:"column:student_answer_3"` // 学生的答案
}

// TableName get sql table name.获取数据库表名
func (m *BfqAnswer) TableName() string {
	return "bfq_answer"
}
