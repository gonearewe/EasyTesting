package models

// Teacher [...]
type Teacher struct {
	ID        int    `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                                    // 用作主键
	TeacherID string `gorm:"unique;index:teacher_id_2;column:teacher_id;type:varchar(10);not null" json:"teacher_id"` // 工号
	Name      string `gorm:"column:name;type:varchar(50);not null" json:"name"`                                       // 姓名
	Password  string `gorm:"column:password;type:varchar(100);not null" json:"-"`                              // 加盐后的密码
	Salt      string `gorm:"column:salt;type:varchar(50);not null" json:"-"`                                       // 盐
	IsAdmin   bool   `gorm:"column:is_admin;type:tinyint(1);not null;default:0" json:"is_admin"`                      // 是否为超级管理员
}

// TableName get sql table name.获取数据库表名
func (m *Teacher) TableName() string {
	return "teacher"
}
