package models

type Tfq struct {
	ID                 int     `gorm:"primaryKey;column:id;type:int(10);not null"`                                        // 用作主键
	PublisherTeacherID int16   `gorm:"index:publisher_teacher_id;column:publisher_teacher_id;type:smallint(10);not null"` // 创建本题的教师的工号
	Teacher            Teacher `gorm:"joinForeignKey:publisher_teacher_id;foreignKey:teacher_id"`
	Stem               string  `gorm:"column:stem;type:tinytext;not null"` // 题干
	Answer             []uint8 `gorm:"column:answer;type:bit(1);not null"` // 0：错，1：对
}
