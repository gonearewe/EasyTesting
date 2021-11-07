package models

type Mcq struct {
	ID                 int     `gorm:"primaryKey;column:id;type:int(10);not null"`                                        // 用作主键
	PublisherTeacherID int16   `gorm:"index:publisher_teacher_id;column:publisher_teacher_id;type:smallint(10);not null"` // 创建本题的教师的工号
	Teacher            Teacher `gorm:"joinForeignKey:publisher_teacher_id;foreignKey:teacher_id"`
	Stem               string  `gorm:"column:stem;type:tinytext;not null"`              // 题干
	Choice1            string  `gorm:"column:choice_1;type:tinytext;not null"`          // 选项的内容
	Choice2            string  `gorm:"column:choice_2;type:tinytext;not null"`          // 选项的内容
	Choice3            string  `gorm:"column:choice_3;type:tinytext;not null"`          // 选项的内容
	Choice4            string  `gorm:"column:choice_4;type:tinytext;not null"`          // 选项的内容
	Choice5            string  `gorm:"column:choice_5;type:tinytext"`                   // 选项的内容
	Choice6            string  `gorm:"column:choice_6;type:tinytext"`                   // 选项的内容
	Choice7            string  `gorm:"column:choice_7;type:tinytext"`                   // 选项的内容
	IsMaq              []uint8 `gorm:"column:is_maq;type:bit(1);not null;default:b'0'"` // 是不是多选题，0：否，1：是
	RightAnswer        string  `gorm:"column:right_answer;type:char(7)"`                // 答案，按升序包含所有正确选项的索引，如 "5"、"124"、"67"
}
