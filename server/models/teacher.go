package models

type Teacher struct {
	ID        int     `gorm:"primaryKey;column:id;type:int(10);not null"`          // 用作主键
	TeacherID int16   `gorm:"unique;column:teacher_id;type:smallint(10);not null"` // 工号
	Name      string  `gorm:"column:name;type:varchar(50);not null"`               // 姓名
	Password  string  `gorm:"column:password;type:varchar(200);not null"`          // 加盐后的密码
	IsAdmin   []uint8 `gorm:"column:is_admin;type:bit(1);not null;default:b'0'"`   // 是否为超级管理员，0：否，1：是
}
