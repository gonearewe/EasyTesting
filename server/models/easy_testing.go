package models

import (
	"time"
)

// Bfq [...]
type Bfq struct {
	ID                 int     `gorm:"primaryKey;column:id;type:int(10);not null"`                                        // 用作主键
	PublisherTeacherID int16   `gorm:"index:publisher_teacher_id;column:publisher_teacher_id;type:smallint(10);not null"` // 创建本题的教师的工号
	Teacher            Teacher `gorm:"joinForeignKey:publisher_teacher_id;foreignKey:teacher_id"`
	Stem               string  `gorm:"column:stem;type:tinytext;not null"`        // 题干
	BlankNum           bool    `gorm:"column:blank_num;type:tinyint(1);not null"` // 要填的空的数目，若大于 1，则说明是 crq
	Answer1            string  `gorm:"column:answer_1;type:tinytext;not null"`    // 填空的答案
	Answer2            string  `gorm:"column:answer_2;type:tinytext"`             // 填空的答案
	Answer3            string  `gorm:"column:answer_3;type:tinytext"`             // 填空的答案
	Answer4            string  `gorm:"column:answer_4;type:tinytext"`             // 填空的答案
	Answer5            string  `gorm:"column:answer_5;type:tinytext"`             // 填空的答案
}

// BfqColumns get sql column name.获取数据库列名
var BfqColumns = struct {
	ID                 string
	PublisherTeacherID string
	Stem               string
	BlankNum           string
	Answer1            string
	Answer2            string
	Answer3            string
	Answer4            string
	Answer5            string
}{
	ID:                 "id",
	PublisherTeacherID: "publisher_teacher_id",
	Stem:               "stem",
	BlankNum:           "blank_num",
	Answer1:            "answer_1",
	Answer2:            "answer_2",
	Answer3:            "answer_3",
	Answer4:            "answer_4",
	Answer5:            "answer_5",
}

// BfqAnswer [...]
type BfqAnswer struct {
	ID             int         `gorm:"primaryKey;column:id;type:int(10);not null"`       // 用作主键
	BfqID          int         `gorm:"index:bfq_id;column:bfq_id;type:int(10);not null"` // 连接 bfq
	Bfq            Bfq         `gorm:"joinForeignKey:bfq_id;foreignKey:id"`
	ExamSessionID  int         `gorm:"index:exam_session_id;column:exam_session_id;type:int(10);not null"` // 连接 exam_session
	ExamSession    ExamSession `gorm:"joinForeignKey:exam_session_id;foreignKey:id"`
	StudentAnswer1 string      `gorm:"column:student_answer_1;type:tinytext"`          // 学生的答案
	StudentAnswer2 string      `gorm:"column:student_answer_2;type:tinytext"`          // 学生的答案
	StudentAnswer3 string      `gorm:"column:student_answer_3;type:tinytext"`          // 学生的答案
	StudentAnswer4 string      `gorm:"column:student_answer_4;type:tinytext"`          // 学生的答案
	StudentAnswer5 string      `gorm:"column:student_answer_5;type:tinytext"`          // 学生的答案
	Score          uint8       `gorm:"column:score;type:tinyint(3) unsigned;not null"` // 本题分值
}

// BfqAnswerColumns get sql column name.获取数据库列名
var BfqAnswerColumns = struct {
	ID             string
	BfqID          string
	ExamSessionID  string
	StudentAnswer1 string
	StudentAnswer2 string
	StudentAnswer3 string
	StudentAnswer4 string
	StudentAnswer5 string
	Score          string
}{
	ID:             "id",
	BfqID:          "bfq_id",
	ExamSessionID:  "exam_session_id",
	StudentAnswer1: "student_answer_1",
	StudentAnswer2: "student_answer_2",
	StudentAnswer3: "student_answer_3",
	StudentAnswer4: "student_answer_4",
	StudentAnswer5: "student_answer_5",
	Score:          "score",
}

// Exam [...]
type Exam struct {
	ID                 int       `gorm:"primaryKey;column:id;type:int(10);not null"`             // 用作主键
	PublisherTeacherID int16     `gorm:"column:publisher_teacher_id;type:smallint(10);not null"` // 发布考试的教师的工号
	StartTime          time.Time `gorm:"column:start_time;type:datetime;not null"`               // 考试开始时间
	EndTime            time.Time `gorm:"column:end_time;type:datetime;not null"`                 // 考试结束时间
	TimeAllowed        string    `gorm:"column:time_allowed;type:varchar(200);not null"`         // 考生答题时间
}

// ExamColumns get sql column name.获取数据库列名
var ExamColumns = struct {
	ID                 string
	PublisherTeacherID string
	StartTime          string
	EndTime            string
	TimeAllowed        string
}{
	ID:                 "id",
	PublisherTeacherID: "publisher_teacher_id",
	StartTime:          "start_time",
	EndTime:            "end_time",
	TimeAllowed:        "time_allowed",
}

// ExamSession [...]
type ExamSession struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(10);not null"`         // 用作主键
	ExamID    int       `gorm:"index:exam_id;column:exam_id;type:int(10);not null"` // 连接 exam
	Exam      Exam      `gorm:"joinForeignKey:exam_id;foreignKey:id"`
	StudentID int16     `gorm:"index:student_id;column:student_id;type:smallint(10);not null"` // 连接 student
	Student   Student   `gorm:"joinForeignKey:student_id;foreignKey:student_id"`
	StartTime time.Time `gorm:"column:start_time;type:datetime;not null"` // 作答开始时间
	EndTime   time.Time `gorm:"column:end_time;type:datetime"`            // 交卷时间
	Score     uint8     `gorm:"column:score;type:tinyint(3) unsigned"`    // 最终成绩
}

// ExamSessionColumns get sql column name.获取数据库列名
var ExamSessionColumns = struct {
	ID        string
	ExamID    string
	StudentID string
	StartTime string
	EndTime   string
	Score     string
}{
	ID:        "id",
	ExamID:    "exam_id",
	StudentID: "student_id",
	StartTime: "start_time",
	EndTime:   "end_time",
	Score:     "score",
}

// Mcq [...]
type Mcq struct {
	ID                 int     `gorm:"primaryKey;column:id;type:int(10);not null"`                                        // 用作主键
	PublisherTeacherID int16   `gorm:"index:publisher_teacher_id;column:publisher_teacher_id;type:smallint(10);not null"` // 创建本题的教师的工号
	Teacher            Teacher `gorm:"joinForeignKey:publisher_teacher_id;foreignKey:teacher_id"`
	Stem               string  `gorm:"column:stem;type:tinytext;not null"`              // 题干
	Choice1            string  `gorm:"column:choice_1;type:tinytext;not null"`          // 选项的内容
	Choice2            string  `gorm:"column:choice_2;type:tinytext;not null"`          // 选项的内容
	Choice3            string  `gorm:"column:choice_3;type:tinytext;not null"`          // 选项的内容
	Choice4            string  `gorm:"column:choice_4;type:tinytext;not null"`          // 选项的内容
	Choice5            string  `gorm:"column:choice_5;type:tinytext"`                   // 选项的内容
	Choice6            string  `gorm:"column:choice_6;type:tinytext"`                   // 选项的内容
	Choice7            string  `gorm:"column:choice_7;type:tinytext"`                   // 选项的内容
	IsMaq              []uint8 `gorm:"column:is_maq;type:bit(1);not null;default:b'0'"` // 是不是多选题，0：否，1：是
	RightAnswer        string  `gorm:"column:right_answer;type:char(7)"`                // 答案，按升序包含所有正确选项的索引，如 "5"、"124"、"67"
}

// McqColumns get sql column name.获取数据库列名
var McqColumns = struct {
	ID                 string
	PublisherTeacherID string
	Stem               string
	Choice1            string
	Choice2            string
	Choice3            string
	Choice4            string
	Choice5            string
	Choice6            string
	Choice7            string
	IsMaq              string
	RightAnswer        string
}{
	ID:                 "id",
	PublisherTeacherID: "publisher_teacher_id",
	Stem:               "stem",
	Choice1:            "choice_1",
	Choice2:            "choice_2",
	Choice3:            "choice_3",
	Choice4:            "choice_4",
	Choice5:            "choice_5",
	Choice6:            "choice_6",
	Choice7:            "choice_7",
	IsMaq:              "is_maq",
	RightAnswer:        "right_answer",
}

// McqAnswer [...]
type McqAnswer struct {
	ID            int         `gorm:"primaryKey;column:id;type:int(10);not null"`       // 用作主键
	McqID         int         `gorm:"index:mcq_id;column:mcq_id;type:int(10);not null"` // 连接 mcq
	Mcq           Mcq         `gorm:"joinForeignKey:mcq_id;foreignKey:id"`
	ExamSessionID int         `gorm:"index:exam_session_id;column:exam_session_id;type:int(10);not null"` // 连接 exam_session
	ExamSession   ExamSession `gorm:"joinForeignKey:exam_session_id;foreignKey:id"`
	RightAnswer   string      `gorm:"column:right_answer;type:char(7);not null"`      // 正确答案，与 mcq 中同名字段保持一致
	StudentAnswer string      `gorm:"column:student_answer;type:char(7)"`             // 学生的答案
	Score         uint8       `gorm:"column:score;type:tinyint(3) unsigned;not null"` // 本题分值
}

// McqAnswerColumns get sql column name.获取数据库列名
var McqAnswerColumns = struct {
	ID            string
	McqID         string
	ExamSessionID string
	RightAnswer   string
	StudentAnswer string
	Score         string
}{
	ID:            "id",
	McqID:         "mcq_id",
	ExamSessionID: "exam_session_id",
	RightAnswer:   "right_answer",
	StudentAnswer: "student_answer",
	Score:         "score",
}

// Student [...]
type Student struct {
	ID        int    `gorm:"primaryKey;column:id;type:int(10);not null"`          // 用作主键
	StudentID int16  `gorm:"unique;column:student_id;type:smallint(10);not null"` // 学号
	Name      string `gorm:"column:name;type:varchar(50);not null"`               // 姓名
	ClassID   int16  `gorm:"column:class_id;type:smallint(10);not null"`          // 班号
}

// StudentColumns get sql column name.获取数据库列名
var StudentColumns = struct {
	ID        string
	StudentID string
	Name      string
	ClassID   string
}{
	ID:        "id",
	StudentID: "student_id",
	Name:      "name",
	ClassID:   "class_id",
}

// Teacher [...]
type Teacher struct {
	ID        int     `gorm:"primaryKey;column:id;type:int(10);not null"`          // 用作主键
	TeacherID int16   `gorm:"unique;column:teacher_id;type:smallint(10);not null"` // 工号
	Name      string  `gorm:"column:name;type:varchar(50);not null"`               // 姓名
	Password  string  `gorm:"column:password;type:varchar(200);not null"`          // 加盐后的密码
	IsAdmin   []uint8 `gorm:"column:is_admin;type:bit(1);not null;default:b'0'"`   // 是否为超级管理员，0：否，1：是
}

// TeacherColumns get sql column name.获取数据库列名
var TeacherColumns = struct {
	ID        string
	TeacherID string
	Name      string
	Password  string
	IsAdmin   string
}{
	ID:        "id",
	TeacherID: "teacher_id",
	Name:      "name",
	Password:  "password",
	IsAdmin:   "is_admin",
}

// Tfq [...]
type Tfq struct {
	ID                 int     `gorm:"primaryKey;column:id;type:int(10);not null"`                                        // 用作主键
	PublisherTeacherID int16   `gorm:"index:publisher_teacher_id;column:publisher_teacher_id;type:smallint(10);not null"` // 创建本题的教师的工号
	Teacher            Teacher `gorm:"joinForeignKey:publisher_teacher_id;foreignKey:teacher_id"`
	Stem               string  `gorm:"column:stem;type:tinytext;not null"` // 题干
	Answer             []uint8 `gorm:"column:answer;type:bit(1);not null"` // 0：错，1：对
}

// TfqColumns get sql column name.获取数据库列名
var TfqColumns = struct {
	ID                 string
	PublisherTeacherID string
	Stem               string
	Answer             string
}{
	ID:                 "id",
	PublisherTeacherID: "publisher_teacher_id",
	Stem:               "stem",
	Answer:             "answer",
}

// TfqAnswer [...]
type TfqAnswer struct {
	ID            int         `gorm:"primaryKey;column:id;type:int(10);not null"`       // 用作主键
	TfqID         int         `gorm:"index:tfq_id;column:tfq_id;type:int(10);not null"` // 连接 tfq
	Tfq           Tfq         `gorm:"joinForeignKey:tfq_id;foreignKey:id"`
	ExamSessionID int         `gorm:"index:exam_session_id;column:exam_session_id;type:int(10);not null"` // 连接 exam_session
	ExamSession   ExamSession `gorm:"joinForeignKey:exam_session_id;foreignKey:id"`
	RightAnswer   []uint8     `gorm:"column:right_answer;type:bit(1);not null"`       // 正确答案，与 tfq 中同名字段保持一致
	StudentAnswer []uint8     `gorm:"column:student_answer;type:bit(1)"`              // 学生的答案，0：错，1：对
	Score         uint8       `gorm:"column:score;type:tinyint(3) unsigned;not null"` // 本题分值
}

// TfqAnswerColumns get sql column name.获取数据库列名
var TfqAnswerColumns = struct {
	ID            string
	TfqID         string
	ExamSessionID string
	RightAnswer   string
	StudentAnswer string
	Score         string
}{
	ID:            "id",
	TfqID:         "tfq_id",
	ExamSessionID: "exam_session_id",
	RightAnswer:   "right_answer",
	StudentAnswer: "student_answer",
	Score:         "score",
}
