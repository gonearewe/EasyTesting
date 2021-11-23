package models

// Teacher [...]
type Teacher struct {
	ID        int    `gorm:"primaryKey;column:id"` // 用作主键
	TeacherID string `gorm:"column:teacher_id"`    // 工号
	Name      string `gorm:"column:name"`          // 姓名
	Password  string `gorm:"column:password"`      // 加盐后的密码
	Salt      string `gorm:"column:salt"`          // 盐
	IsAdmin   bool   `gorm:"column:is_admin"`      // 是否为超级管理员
}

// TableName get sql table name.获取数据库表名
func (m *Teacher) TableName() string {
	return "teacher"
}
