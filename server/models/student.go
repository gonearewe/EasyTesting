package models

// Student [...]
type Student struct {
	ID        int    `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                 // 用作主键
	StudentID string `gorm:"unique;column:student_id;type:varchar(10);not null" json:"student_id"` // 学号
	Name      string `gorm:"column:name;type:varchar(50);not null" json:"name"`                    // 姓名
	ClassID   string `gorm:"column:class_id;type:varchar(10);not null" json:"class_id"`            // 班号
}

// TableName get sql table name.获取数据库表名
func (m *Student) TableName() string {
	return "student"
}
