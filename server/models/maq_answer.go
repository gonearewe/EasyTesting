package models

// MaqAnswer [...]
type MaqAnswer struct {
	ID            int    `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                                      // 用作主键
	MaqID         int    `gorm:"index:maq_id;column:maq_id;type:int(10);not null" json:"maq_id"`                            // 连接 maq
	ExamSessionID int    `gorm:"index:exam_session_id;column:exam_session_id;type:int(10);not null" json:"exam_session_id"` // 连接 exam_session
	RightAnswer   string `gorm:"column:right_answer;type:char(7);not null" json:"right_answer"`                             // 正确答案，与 maq 中同名字段保持一致
	StudentAnswer string `gorm:"column:student_answer;type:char(7)" json:"student_answer"`                                  // 学生的答案
}

// TableName get sql table name.获取数据库表名
func (m *MaqAnswer) TableName() string {
	return "maq_answer"
}