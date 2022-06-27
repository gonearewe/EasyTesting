package models

import (
	"time"
)

// Crq [...]
type Crq struct {
	ID                  int       `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                                                         // 用作主键
	PublisherTeacherID  string    `gorm:"index:publisher_teacher_id;column:publisher_teacher_id;type:varchar(10);not null" json:"publisher_teacher_id"` // 创建本题的教师的工号
	Stem                string    `gorm:"column:stem;type:text;not null" json:"stem"`                                                                   // 题干
	BlankNum            int8      `gorm:"column:blank_num;type:tinyint(2);not null" json:"blank_num"`                                                   // 要填的空的数目
	Answer1             string    `gorm:"column:answer_1;type:tinytext;not null" json:"answer_1"`                                                       // 填空的答案
	Answer2             string    `gorm:"column:answer_2;type:tinytext;not null" json:"answer_2"`                                                       // 填空的答案
	Answer3             string    `gorm:"column:answer_3;type:tinytext" json:"answer_3"`                                                                // 填空的答案
	Answer4             string    `gorm:"column:answer_4;type:tinytext" json:"answer_4"`                                                                // 填空的答案
	Answer5             string    `gorm:"column:answer_5;type:tinytext" json:"answer_5"`                                                                // 填空的答案
	Answer6             string    `gorm:"column:answer_6;type:tinytext" json:"answer_6"`                                                                // 填空的答案
	OverallCorrectScore int       `gorm:"column:overall_correct_score;type:int(11);not null;default:0" json:"overall_correct_score"`                    // 此题在所有出现中的总得分数*10，即保存到小数点后一位
	OverallScore        int       `gorm:"column:overall_score;type:int(11);not null;default:0" json:"overall_score"`                                    // 此题在所有出现中的总分数*10，即保存到小数点后一位
	CreatedAt           time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"-"`
	LastUpdatedAt       time.Time `gorm:"column:last_updated_at;type:datetime" json:"-"`
}

// TableName get sql table name.获取数据库表名
func (m *Crq) TableName() string {
	return "crq"
}
