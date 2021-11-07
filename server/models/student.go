package models

type Student struct {
	ID        int    `gorm:"primaryKey;column:id;type:int(10);not null"`          // 用作主键
	StudentID int16  `gorm:"unique;column:student_id;type:smallint(10);not null"` // 学号
	Name      string `gorm:"column:name;type:varchar(50);not null"`               // 姓名
	ClassID   int16  `gorm:"column:class_id;type:smallint(10);not null"`          // 班号
}
