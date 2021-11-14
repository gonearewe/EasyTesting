package models

type Teacher struct {
	ID        int     `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                 // 用作主键
	TeacherID string  `gorm:"unique;column:teacher_id;type:varchar(10);not null" json:"teacher_id"` // 工号
	Name      string  `gorm:"column:name;type:varchar(50);not null" json:"name"`                    // 姓名
	Password  string  `gorm:"column:password;type:varchar(200);not null" json:"password"`           // 加盐后的密码
	Salt      string  `gorm:"column:salt;type:varchar(200);not null"`                               // 盐
	IsAdmin   []uint8 `gorm:"column:is_admin;type:bit(1);not null;default:b'0'" json:"is_admin"`    // 是否为超级管理员，0：否，1：是
}

func (Teacher) TableName() string {
	return "teacher"
}
