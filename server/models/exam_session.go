package models

import "time"

type ExamSession struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(10);not null"`         // 用作主键
	ExamID    int       `gorm:"index:exam_id;column:exam_id;type:int(10);not null"` // 连接 exam
	Exam      Exam      `gorm:"joinForeignKey:exam_id;foreignKey:id"`
	StudentID int16     `gorm:"index:student_id;column:student_id;type:smallint(10);not null"` // 连接 student
	Student   Student   `gorm:"joinForeignKey:student_id;foreignKey:student_id"`
	StartTime time.Time `gorm:"column:start_time;type:datetime;not null"` // 作答开始时间
	EndTime   time.Time `gorm:"column:end_time;type:datetime"`            // 交卷时间
	Score     uint8     `gorm:"column:score;type:tinyint(3) unsigned"`    // 最终成绩
}
