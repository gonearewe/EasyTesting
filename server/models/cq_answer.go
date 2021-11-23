package models

// CqAnswer [...]
type CqAnswer struct {
	ID            int         `gorm:"primaryKey;column:id"` // 用作主键
	CqID          int         `gorm:"column:cq_id"`         // 连接 cq
	Cq            Cq          `gorm:"joinForeignKey:cq_id;foreignKey:id"`
	ExamSessionID int         `gorm:"column:exam_session_id"` // 连接 exam_session
	ExamSession   ExamSession `gorm:"joinForeignKey:exam_session_id;foreignKey:id"`
	StudentAnswer string      `gorm:"column:student_answer"` // 学生的答案
}

// TableName get sql table name.获取数据库表名
func (m *CqAnswer) TableName() string {
	return "cq_answer"
}
