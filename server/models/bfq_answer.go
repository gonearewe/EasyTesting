package models

type BfqAnswer struct {
	ID             int         `gorm:"primaryKey;column:id;type:int(10);not null"`       // 用作主键
	BfqID          int         `gorm:"index:bfq_id;column:bfq_id;type:int(10);not null"` // 连接 bfq
	Bfq            Bfq         `gorm:"joinForeignKey:bfq_id;foreignKey:id"`
	ExamSessionID  int         `gorm:"index:exam_session_id;column:exam_session_id;type:int(10);not null"` // 连接 exam_session
	ExamSession    ExamSession `gorm:"joinForeignKey:exam_session_id;foreignKey:id"`
	StudentAnswer1 string      `gorm:"column:student_answer_1;type:tinytext"`          // 学生的答案
	StudentAnswer2 string      `gorm:"column:student_answer_2;type:tinytext"`          // 学生的答案
	StudentAnswer3 string      `gorm:"column:student_answer_3;type:tinytext"`          // 学生的答案
	StudentAnswer4 string      `gorm:"column:student_answer_4;type:tinytext"`          // 学生的答案
	StudentAnswer5 string      `gorm:"column:student_answer_5;type:tinytext"`          // 学生的答案
	Score          uint8       `gorm:"column:score;type:tinyint(3) unsigned;not null"` // 本题分值
}
