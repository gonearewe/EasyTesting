package models

import (
	"time"
)

// ExamSession [...]
type ExamSession struct {
	ID            int       `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                                               // 用作主键
	ExamID        int       `gorm:"uniqueIndex:exam_id;column:exam_id;type:int(10);not null" json:"-"`                                  // 连接 exam
	StudentID     string    `gorm:"uniqueIndex:exam_id;index:student_id;column:student_id;type:varchar(10);not null" json:"student_id"` // 连接 student
	StudentName   string    `gorm:"column:student_name;type:varchar(50);not null" json:"student_name"`                                  // 考生的姓名
	StartTime     time.Time `gorm:"column:start_time;type:datetime;not null" json:"start_time"`                                         // 作答开始时刻
	TimeAllowed   int8      `gorm:"column:time_allowed;type:tinyint(3);not null" json:"time_allowed"`                                   // 考生答题时间，单位：分钟
	EndTime       time.Time `gorm:"column:end_time;type:datetime" json:"end_time"`                                                      // 交卷时刻
	Score         int16     `gorm:"column:score;type:smallint(6);not null" json:"score"`                                                // 最终成绩*10，即保存到小数点后一位
	CreatedAt     time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	LastUpdatedAt time.Time `gorm:"column:last_updated_at;type:datetime" json:"last_updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *ExamSession) TableName() string {
	return "exam_session"
}
