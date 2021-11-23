package models

// Student [...]
type Student struct {
	ID        int    `gorm:"primaryKey;column:id"` // 用作主键
	StudentID string `gorm:"column:student_id"`    // 学号
	Name      string `gorm:"column:name"`          // 姓名
	ClassID   string `gorm:"column:class_id"`      // 班号
}

// TableName get sql table name.获取数据库表名
func (m *Student) TableName() string {
	return "student"
}
