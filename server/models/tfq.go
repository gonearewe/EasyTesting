package models

// Tfq [...]
type Tfq struct {
	ID                 int    `gorm:"primaryKey;column:id"`        // 用作主键
	PublisherTeacherID string `gorm:"column:publisher_teacher_id"` // 创建本题的教师的工号
	Stem               string `gorm:"column:stem"`                 // 题干
	Answer             bool   `gorm:"column:answer"`               // 正确答案
}

// TableName get sql table name.获取数据库表名
func (m *Tfq) TableName() string {
	return "tfq"
}
