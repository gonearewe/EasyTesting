package models

import "time"

type Exam struct {
	ID                 int       `gorm:"primaryKey;column:id;type:int(10);not null"`             // 用作主键
	PublisherTeacherID int16     `gorm:"column:publisher_teacher_id;type:smallint(10);not null"` // 发布考试的教师的工号
	StartTime          time.Time `gorm:"column:start_time;type:datetime;not null"`               // 考试开始时间
	EndTime            time.Time `gorm:"column:end_time;type:datetime;not null"`                 // 考试结束时间
	TimeAllowed        string    `gorm:"column:time_allowed;type:varchar(200);not null"`         // 考生答题时间
}
