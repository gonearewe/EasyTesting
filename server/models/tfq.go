package models

// Tfq [...]
type Tfq struct {
	ID                 int    `gorm:"primaryKey;column:id;type:int(10);not null" json:"id"`                                                         // 用作主键
	PublisherTeacherID string `gorm:"index:publisher_teacher_id;column:publisher_teacher_id;type:varchar(10);not null" json:"publisher_teacher_id"` // 创建本题的教师的工号
	Stem               string `gorm:"column:stem;type:text;not null" json:"stem"`                                                                   // 题干
	Answer             bool   `gorm:"column:answer;type:tinyint(1);not null" json:"answer"`                                                                // 填空的答案
	OverallCorrectScore int   `gorm:"column:overall_correct_score;type:int;not null" json:"overall_correct_score"` // 此题在所有出现中的总得分数*10，即保存到小数点后一位
	OverallScore int   `gorm:"column:overall_score;type:int;not null" json:"overall_score"` // 此题在所有出现中的总分数*10，即保存到小数点后一位                                                         // 正确答案
}

// TableName get sql table name.获取数据库表名
func (m *Tfq) TableName() string {
	return "tfq"
}
