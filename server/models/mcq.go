package models

// Mcq [...]
type Mcq struct {
	ID                 int    `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                                                         // 用作主键
	PublisherTeacherID string `gorm:"index:publisher_teacher_id;column:publisher_teacher_id;type:varchar(10);not null" json:"publisher_teacher_id"` // 创建本题的教师的工号
	Stem               string `gorm:"column:stem;type:text;not null" json:"stem"`                                                                   // 题干
	Choice1            string `gorm:"column:choice_1;type:text;not null" json:"choice_1"`                                                           // 选项的内容
	Choice2            string `gorm:"column:choice_2;type:text;not null" json:"choice_2"`                                                           // 选项的内容
	Choice3            string `gorm:"column:choice_3;type:text;not null" json:"choice_3"`                                                           // 选项的内容
	Choice4            string `gorm:"column:choice_4;type:text;not null" json:"choice_4"`                                                           // 选项的内容
	RightAnswer        string `gorm:"column:right_answer;type:char(1);not null" json:"right_answer"`                                                                // 填空的答案
	OverallCorrectScore int   `gorm:"column:overall_correct_score;type:int;not null" json:"overall_correct_score"` // 此题在所有出现中的总得分数*10，即保存到小数点后一位
	OverallScore int   `gorm:"column:overall_score;type:int;not null" json:"overall_score"` // 此题在所有出现中的总分数*10，即保存到小数点后一位                                                // 答案，正确选项的索引，如 "4"、"1"
}

// TableName get sql table name.获取数据库表名
func (m *Mcq) TableName() string {
	return "mcq"
}
