package models

// Maq [...]
type Maq struct {
	ID                 int    `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                                                         // 用作主键
	PublisherTeacherID string `gorm:"index:publisher_teacher_id;column:publisher_teacher_id;type:varchar(10);not null" json:"publisher_teacher_id"` // 创建本题的教师的工号
	Stem               string `gorm:"column:stem;type:text;not null" json:"stem"`                                                                   // 题干
	Choice1            string `gorm:"column:choice_1;type:text;not null" json:"choice_1"`                                                           // 选项的内容
	Choice2            string `gorm:"column:choice_2;type:text;not null" json:"choice_2"`                                                           // 选项的内容
	Choice3            string `gorm:"column:choice_3;type:text;not null" json:"choice_3"`                                                           // 选项的内容
	Choice4            string `gorm:"column:choice_4;type:text;not null" json:"choice_4"`                                                           // 选项的内容
	Choice5            string `gorm:"column:choice_5;type:text" json:"choice_5"`                                                                    // 选项的内容
	Choice6            string `gorm:"column:choice_6;type:text" json:"choice_6"`                                                                    // 选项的内容
	Choice7            string `gorm:"column:choice_7;type:text" json:"choice_7"`                                                                    // 选项的内容
	RightAnswer        string `gorm:"column:right_answer;type:char(7);not null" json:"right_answer"`                                                // 答案，按升序包含所有正确选项的索引，如 "5"、"124"、"67"
}

// TableName get sql table name.获取数据库表名
func (m *Maq) TableName() string {
	return "maq"
}
