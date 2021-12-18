package models

// Tfq [...]
type Tfq struct {
	ID                 int    `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                                                         // 用作主键
	PublisherTeacherID string `gorm:"index:publisher_teacher_id;column:publisher_teacher_id;type:varchar(10);not null" json:"publisher_teacher_id"` // 创建本题的教师的工号
	Stem               string `gorm:"column:stem;type:text;not null" json:"stem"`                                                                   // 题干
	Answer             bool   `gorm:"column:answer;type:tinyint(1);not null" json:"answer"`                                                         // 正确答案
}

// TableName get sql table name.获取数据库表名
func (m *Tfq) TableName() string {
	return "tfq"
}
