package models

import (
	"time"
)

// Cq [...]
type Cq struct {
	ID                  int       `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                                                         // 用作主键
	PublisherTeacherID  string    `gorm:"index:publisher_teacher_id;column:publisher_teacher_id;type:varchar(10);not null" json:"publisher_teacher_id"` // 创建本题的教师的工号
	Stem                string    `gorm:"column:stem;type:text;not null" json:"stem"`                                                                   // 题干
	IsInputFromFile     bool      `gorm:"column:is_input_from_file;type:tinyint(1);not null" json:"is_input_from_file"`                                 // 程序输入是否从文件读取，若为否，从命令行读取
	IsOutputToFile      bool      `gorm:"column:is_output_to_file;type:tinyint(1);not null" json:"is_output_to_file"`                                   // 程序输出是否写入文件，若为否，输出到命令行
	Input               string    `gorm:"column:input;type:text;not null" json:"input"`                                                                 // 程序的输入
	Output              string    `gorm:"column:output;type:text;not null" json:"output"`                                                               // 程序的输出
	Template            string    `gorm:"column:template;type:text;not null" json:"template"`                                                           // 题目的初始模板
	OverallCorrectScore int       `gorm:"column:overall_correct_score;type:int(11);not null;default:0" json:"overall_correct_score"`                    // 此题在所有出现中的总得分数*10，即保存到小数点后一位
	OverallScore        int       `gorm:"column:overall_score;type:int(11);not null;default:0" json:"overall_score"`                                    // 此题在所有出现中的总分数*10，即保存到小数点后一位
	CreatedAt           time.Time `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	LastUpdatedAt       time.Time `gorm:"column:last_updated_at;type:datetime" json:"last_updated_at"`
}

// TableName get sql table name.获取数据库表名
func (m *Cq) TableName() string {
	return "cq"
}
