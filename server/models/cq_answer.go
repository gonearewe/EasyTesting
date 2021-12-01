package models

// CqAnswer [...]
type CqAnswer struct {
	ID            int    `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                                      // 用作主键
	CqID          int    `gorm:"index:cq_id;column:cq_id;type:int(10);not null" json:"cq_id"`                               // 连接 cq
	ExamSessionID int    `gorm:"index:exam_session_id;column:exam_session_id;type:int(10);not null" json:"exam_session_id"` // 连接 exam_session
	StudentAnswer string `gorm:"column:student_answer;type:tinytext" json:"student_answer"`                                 // 学生的答案
}

// TableName get sql table name.获取数据库表名
func (m *CqAnswer) TableName() string {
	return "cq_answer"
}
