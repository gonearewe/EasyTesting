package models

// Mcq [...]
type Mcq struct {
	ID                 int    `gorm:"primaryKey;column:id"`        // 用作主键
	PublisherTeacherID string `gorm:"column:publisher_teacher_id"` // 创建本题的教师的工号
	Stem               string `gorm:"column:stem"`                 // 题干
	Choice1            string `gorm:"column:choice_1"`             // 选项的内容
	Choice2            string `gorm:"column:choice_2"`             // 选项的内容
	Choice3            string `gorm:"column:choice_3"`             // 选项的内容
	Choice4            string `gorm:"column:choice_4"`             // 选项的内容
	RightAnswer        string `gorm:"column:right_answer"`         // 答案，正确选项的索引，如 "4"、"1"
}

// TableName get sql table name.获取数据库表名
func (m *Mcq) TableName() string {
	return "mcq"
}
