package models

type TfqAnswer struct {
	ID            int         `gorm:"primaryKey;column:id;type:int(10);not null"`       // 用作主键
	TfqID         int         `gorm:"index:tfq_id;column:tfq_id;type:int(10);not null"` // 连接 tfq
	Tfq           Tfq         `gorm:"joinForeignKey:tfq_id;foreignKey:id"`
	ExamSessionID int         `gorm:"index:exam_session_id;column:exam_session_id;type:int(10);not null"` // 连接 exam_session
	ExamSession   ExamSession `gorm:"joinForeignKey:exam_session_id;foreignKey:id"`
	RightAnswer   []uint8     `gorm:"column:right_answer;type:bit(1);not null"`       // 正确答案，与 tfq 中同名字段保持一致
	StudentAnswer []uint8     `gorm:"column:student_answer;type:bit(1)"`              // 学生的答案，0：错，1：对
	Score         uint8       `gorm:"column:score;type:tinyint(3) unsigned;not null"` // 本题分值
}
