package models

// Bfq [...]
type Bfq struct {
	ID                 int    `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                                                         // 用作主键
	PublisherTeacherID string `gorm:"index:publisher_teacher_id;column:publisher_teacher_id;type:varchar(10);not null" json:"publisher_teacher_id"` // 创建本题的教师的工号
	Stem               string `gorm:"column:stem;type:tinytext;not null" json:"stem"`                                                               // 题干
	BlankNum           bool   `gorm:"column:blank_num;type:tinyint(1);not null" json:"blank_num"`                                                   // 要填的空的数目
	Answer1            string `gorm:"column:answer_1;type:tinytext;not null" json:"answer_1"`                                                       // 填空的答案
	Answer2            string `gorm:"column:answer_2;type:tinytext" json:"answer_2"`                                                                // 填空的答案
	Answer3            string `gorm:"column:answer_3;type:tinytext" json:"answer_3"`                                                                // 填空的答案
}

// TableName get sql table name.获取数据库表名
func (m *Bfq) TableName() string {
	return "bfq"
}
