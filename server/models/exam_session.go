package models

import (
	"time"
)

// ExamSession [...]
type ExamSession struct {
	ID          int       `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                           // 用作主键
	ExamID      int       `gorm:"index:exam_id;column:exam_id;type:int(10);not null" json:"exam_id"`              // 连接 exam
	StudentID   string    `gorm:"index:student_id;column:student_id;type:varchar(10);not null" json:"student_id"` // 连接 student
	StartTime   time.Time `gorm:"column:start_time;type:datetime;not null" json:"start_time"`                     // 作答开始时间
	EndTime     time.Time `gorm:"column:end_time;type:datetime" json:"end_time"`                                  // 交卷时间
	AnswerSheet []byte    `gorm:"column:answer_sheet;type:mediumblob" json:"answer_sheet"`                        // 包括考试试题与作答情况的pdf，用于存档
	Score       *uint8    `gorm:"column:score;type:tinyint(3) unsigned" json:"score"`                             // 最终成绩
}

// TableName get sql table name.获取数据库表名
func (m *ExamSession) TableName() string {
	return "exam_session"
}
