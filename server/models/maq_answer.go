package models

// MaqAnswer [...]
type MaqAnswer struct {
	ID            int         `gorm:"primaryKey;column:id"` // 用作主键
	MaqID         int         `gorm:"column:maq_id"`        // 连接 maq
	Maq           Maq         `gorm:"joinForeignKey:maq_id;foreignKey:id"`
	ExamSessionID int         `gorm:"column:exam_session_id"` // 连接 exam_session
	ExamSession   ExamSession `gorm:"joinForeignKey:exam_session_id;foreignKey:id"`
	RightAnswer   string      `gorm:"column:right_answer"`   // 正确答案，与 maq 中同名字段保持一致
	StudentAnswer string      `gorm:"column:student_answer"` // 学生的答案
}

// TableName get sql table name.获取数据库表名
func (m *MaqAnswer) TableName() string {
	return "maq_answer"
}
