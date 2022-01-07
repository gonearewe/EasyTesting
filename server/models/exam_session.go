package models

import (
	"time"
)

// ExamSession [...]
type ExamSession struct {
	ID          int       `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                           // 用作主键
	ExamID      int       `gorm:"index:exam_id;column:exam_id;type:int(10);not null" json:"-"`                    // 连接 exam
	StudentID   string    `gorm:"index:student_id;column:student_id;type:varchar(10);not null" json:"student_id"` // 连接 student
	StudentName string    `gorm:"column:student_name;type:varchar(50);not null" json:"student_name"`
	StartTime   time.Time `gorm:"column:start_time;type:datetime;not null" json:"start_time"` // 作答开始时刻
	EndTime     time.Time `gorm:"column:end_time;type:datetime" json:"end_time"`              // 交卷时刻
	AnswerSheet []byte    `gorm:"column:answer_sheet;type:mediumblob" json:"-"`               // 包括考试试题与作答情况的pdf，用于存档
	Score       *int16    `gorm:"column:score;type:smallint" json:"score"`                    // 最终成绩*10，即保存到小数点后一位
}

// TableName get sql table name.获取数据库表名
func (m *ExamSession) TableName() string {
	return "exam_session"
}
