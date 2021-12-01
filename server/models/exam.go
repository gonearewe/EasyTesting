package models

import (
	"time"
)

// Exam [...]
type Exam struct {
	ID                 int       `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                                                         // 用作主键
	PublisherTeacherID string    `gorm:"index:publisher_teacher_id;column:publisher_teacher_id;type:varchar(10);not null" json:"publisher_teacher_id"` // 发布考试的教师的工号
	StartTime          time.Time `gorm:"column:start_time;type:datetime;not null" json:"start_time"`                                                   // 考试开始时间
	EndTime            time.Time `gorm:"column:end_time;type:datetime;not null" json:"end_time"`                                                       // 考试结束时间
	TimeAllowed        string    `gorm:"column:time_allowed;type:varchar(200);not null" json:"time_allowed"`                                           // 考生答题时间
	McqScore           bool      `gorm:"column:mcq_score;type:tinyint(1) unsigned;not null" json:"mcq_score"`                                          // 单选题每题分数
	McqNum             uint8     `gorm:"column:mcq_num;type:tinyint(2) unsigned;not null" json:"mcq_num"`                                              // 单选题题数
	MaqScore           bool      `gorm:"column:maq_score;type:tinyint(1) unsigned;not null" json:"maq_score"`                                          // 多选题每题分数
	MaqNum             uint8     `gorm:"column:maq_num;type:tinyint(2) unsigned;not null" json:"maq_num"`                                              // 多选题题数
	BfqScore           bool      `gorm:"column:bfq_score;type:tinyint(1) unsigned;not null" json:"bfq_score"`                                          // 填空题每题分数
	BfqNum             uint8     `gorm:"column:bfq_num;type:tinyint(2) unsigned;not null" json:"bfq_num"`                                              // 填空题题数
	TfqScore           bool      `gorm:"column:tfq_score;type:tinyint(1) unsigned;not null" json:"tfq_score"`                                          // 判断题每题分数
	TfqNum             uint8     `gorm:"column:tfq_num;type:tinyint(2) unsigned;not null" json:"tfq_num"`                                              // 判断题题数
	CrqScore           bool      `gorm:"column:crq_score;type:tinyint(1) unsigned;not null" json:"crq_score"`                                          // 代码阅读题每题分数
	CrqNum             bool      `gorm:"column:crq_num;type:tinyint(1) unsigned;not null" json:"crq_num"`                                              // 代码阅读题题数
	CqScore            uint8     `gorm:"column:cq_score;type:tinyint(2) unsigned;not null" json:"cq_score"`                                            // 写代码题每题分数
	CqNum              bool      `gorm:"column:cq_num;type:tinyint(1) unsigned;not null" json:"cq_num"`                                                // 写代码题题数
}

// TableName get sql table name.获取数据库表名
func (m *Exam) TableName() string {
	return "exam"
}
