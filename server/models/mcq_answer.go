package models

type McqAnswer struct {
	ID            int         `gorm:"primaryKey;column:id;type:int(10);not null"`       // 用作主键
	McqID         int         `gorm:"index:mcq_id;column:mcq_id;type:int(10);not null"` // 连接 mcq
	Mcq           Mcq         `gorm:"joinForeignKey:mcq_id;foreignKey:id"`
	ExamSessionID int         `gorm:"index:exam_session_id;column:exam_session_id;type:int(10);not null"` // 连接 exam_session
	ExamSession   ExamSession `gorm:"joinForeignKey:exam_session_id;foreignKey:id"`
	RightAnswer   string      `gorm:"column:right_answer;type:char(7);not null"`      // 正确答案，与 mcq 中同名字段保持一致
	StudentAnswer string      `gorm:"column:student_answer;type:char(7)"`             // 学生的答案
	Score         uint8       `gorm:"column:score;type:tinyint(3) unsigned;not null"` // 本题分值
}
