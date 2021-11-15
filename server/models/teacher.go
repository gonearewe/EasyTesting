package models

type Teacher struct {
	ID        int    `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                 // 用作主键
	TeacherID string `gorm:"unique;column:teacher_id;type:varchar(10);not null" json:"teacher_id"` // 工号
	Name      string `gorm:"column:name;type:varchar(50);not null" json:"name,omitempty"`          // 姓名
	Password  string `gorm:"column:password;type:varchar(200);not null" json:"password,omitempty"` // 加盐后的密码
	Salt      string `gorm:"column:salt;type:varchar(200);not null" json:"-"`                      // 盐
	IsAdmin   bool   `gorm:"column:is_admin;type:bool;not null;default:false" json:"is_admin"`     // 是否为超级管理员
}

func (Teacher) TableName() string {
	return "teacher"
}
