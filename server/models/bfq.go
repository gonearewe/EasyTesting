package models

type Bfq struct {
	ID                 int     `gorm:"primaryKey;column:id;type:int(10);not null"`                                        // 用作主键
	PublisherTeacherID int16   `gorm:"index:publisher_teacher_id;column:publisher_teacher_id;type:smallint(10);not null"` // 创建本题的教师的工号
	Teacher            Teacher `gorm:"joinForeignKey:publisher_teacher_id;foreignKey:teacher_id"`
	Stem               string  `gorm:"column:stem;type:tinytext;not null"`        // 题干
	BlankNum           bool    `gorm:"column:blank_num;type:tinyint(1);not null"` // 要填的空的数目，若大于 1，则说明是 crq
	Answer1            string  `gorm:"column:answer_1;type:tinytext;not null"`    // 填空的答案
	Answer2            string  `gorm:"column:answer_2;type:tinytext"`             // 填空的答案
	Answer3            string  `gorm:"column:answer_3;type:tinytext"`             // 填空的答案
	Answer4            string  `gorm:"column:answer_4;type:tinytext"`             // 填空的答案
	Answer5            string  `gorm:"column:answer_5;type:tinytext"`             // 填空的答案
}
