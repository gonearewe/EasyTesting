package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _BfqAnswerMgr struct {
	*_BaseMgr
}

// BfqAnswerMgr open func
func BfqAnswerMgr(db *gorm.DB) *_BfqAnswerMgr {
	if db == nil {
		panic(fmt.Errorf("BfqAnswerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BfqAnswerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("bfq_answer"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BfqAnswerMgr) GetTableName() string {
	return "bfq_answer"
}

// Reset 重置gorm会话
func (obj *_BfqAnswerMgr) Reset() *_BfqAnswerMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_BfqAnswerMgr) Get() (result BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("bfq").Where("id = ?", result.BfqID).Find(&result.Bfq).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("exam_session").Where("id = ?", result.ExamSessionID).Find(&result.ExamSession).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_BfqAnswerMgr) Gets() (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_BfqAnswerMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Count(count)
}

// ////////////////////////////////////////////////////////////////////////////////

// ////////////////////////option case ////////////////////////////////////////////

// WithID id获取 用作主键
func (obj *_BfqAnswerMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithBfqID bfq_id获取 连接 bfq
func (obj *_BfqAnswerMgr) WithBfqID(bfqID int) Option {
	return optionFunc(func(o *options) { o.query["bfq_id"] = bfqID })
}

// WithExamSessionID exam_session_id获取 连接 exam_session
func (obj *_BfqAnswerMgr) WithExamSessionID(examSessionID int) Option {
	return optionFunc(func(o *options) { o.query["exam_session_id"] = examSessionID })
}

// WithStudentAnswer1 student_answer_1获取 学生的答案
func (obj *_BfqAnswerMgr) WithStudentAnswer1(studentAnswer1 string) Option {
	return optionFunc(func(o *options) { o.query["student_answer_1"] = studentAnswer1 })
}

// WithStudentAnswer2 student_answer_2获取 学生的答案
func (obj *_BfqAnswerMgr) WithStudentAnswer2(studentAnswer2 string) Option {
	return optionFunc(func(o *options) { o.query["student_answer_2"] = studentAnswer2 })
}

// WithStudentAnswer3 student_answer_3获取 学生的答案
func (obj *_BfqAnswerMgr) WithStudentAnswer3(studentAnswer3 string) Option {
	return optionFunc(func(o *options) { o.query["student_answer_3"] = studentAnswer3 })
}

// WithStudentAnswer4 student_answer_4获取 学生的答案
func (obj *_BfqAnswerMgr) WithStudentAnswer4(studentAnswer4 string) Option {
	return optionFunc(func(o *options) { o.query["student_answer_4"] = studentAnswer4 })
}

// WithStudentAnswer5 student_answer_5获取 学生的答案
func (obj *_BfqAnswerMgr) WithStudentAnswer5(studentAnswer5 string) Option {
	return optionFunc(func(o *options) { o.query["student_answer_5"] = studentAnswer5 })
}

// WithScore score获取 本题分值
func (obj *_BfqAnswerMgr) WithScore(score uint8) Option {
	return optionFunc(func(o *options) { o.query["score"] = score })
}

// GetByOption 功能选项模式获取
func (obj *_BfqAnswerMgr) GetByOption(opts ...Option) (result BfqAnswer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("bfq").Where("id = ?", result.BfqID).Find(&result.Bfq).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("exam_session").Where("id = ?", result.ExamSessionID).Find(&result.ExamSession).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_BfqAnswerMgr) GetByOptions(opts ...Option) (results []*BfqAnswer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// ////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 用作主键
func (obj *_BfqAnswerMgr) GetFromID(id int) (result BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("bfq").Where("id = ?", result.BfqID).Find(&result.Bfq).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("exam_session").Where("id = ?", result.ExamSessionID).Find(&result.ExamSession).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找 用作主键
func (obj *_BfqAnswerMgr) GetBatchFromID(ids []int) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromBfqID 通过bfq_id获取内容 连接 bfq
func (obj *_BfqAnswerMgr) GetFromBfqID(bfqID int) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`bfq_id` = ?", bfqID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromBfqID 批量查找 连接 bfq
func (obj *_BfqAnswerMgr) GetBatchFromBfqID(bfqIDs []int) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`bfq_id` IN (?)", bfqIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromExamSessionID 通过exam_session_id获取内容 连接 exam_session
func (obj *_BfqAnswerMgr) GetFromExamSessionID(examSessionID int) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`exam_session_id` = ?", examSessionID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromExamSessionID 批量查找 连接 exam_session
func (obj *_BfqAnswerMgr) GetBatchFromExamSessionID(examSessionIDs []int) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`exam_session_id` IN (?)", examSessionIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStudentAnswer1 通过student_answer_1获取内容 学生的答案
func (obj *_BfqAnswerMgr) GetFromStudentAnswer1(studentAnswer1 string) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`student_answer_1` = ?", studentAnswer1).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStudentAnswer1 批量查找 学生的答案
func (obj *_BfqAnswerMgr) GetBatchFromStudentAnswer1(studentAnswer1s []string) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`student_answer_1` IN (?)", studentAnswer1s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStudentAnswer2 通过student_answer_2获取内容 学生的答案
func (obj *_BfqAnswerMgr) GetFromStudentAnswer2(studentAnswer2 string) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`student_answer_2` = ?", studentAnswer2).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStudentAnswer2 批量查找 学生的答案
func (obj *_BfqAnswerMgr) GetBatchFromStudentAnswer2(studentAnswer2s []string) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`student_answer_2` IN (?)", studentAnswer2s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStudentAnswer3 通过student_answer_3获取内容 学生的答案
func (obj *_BfqAnswerMgr) GetFromStudentAnswer3(studentAnswer3 string) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`student_answer_3` = ?", studentAnswer3).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStudentAnswer3 批量查找 学生的答案
func (obj *_BfqAnswerMgr) GetBatchFromStudentAnswer3(studentAnswer3s []string) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`student_answer_3` IN (?)", studentAnswer3s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStudentAnswer4 通过student_answer_4获取内容 学生的答案
func (obj *_BfqAnswerMgr) GetFromStudentAnswer4(studentAnswer4 string) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`student_answer_4` = ?", studentAnswer4).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStudentAnswer4 批量查找 学生的答案
func (obj *_BfqAnswerMgr) GetBatchFromStudentAnswer4(studentAnswer4s []string) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`student_answer_4` IN (?)", studentAnswer4s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStudentAnswer5 通过student_answer_5获取内容 学生的答案
func (obj *_BfqAnswerMgr) GetFromStudentAnswer5(studentAnswer5 string) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`student_answer_5` = ?", studentAnswer5).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStudentAnswer5 批量查找 学生的答案
func (obj *_BfqAnswerMgr) GetBatchFromStudentAnswer5(studentAnswer5s []string) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`student_answer_5` IN (?)", studentAnswer5s).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromScore 通过score获取内容 本题分值
func (obj *_BfqAnswerMgr) GetFromScore(score uint8) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`score` = ?", score).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromScore 批量查找 本题分值
func (obj *_BfqAnswerMgr) GetBatchFromScore(scores []uint8) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`score` IN (?)", scores).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// ////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_BfqAnswerMgr) FetchByPrimaryKey(id int) (result BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("bfq").Where("id = ?", result.BfqID).Find(&result.Bfq).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("exam_session").Where("id = ?", result.ExamSessionID).Find(&result.ExamSession).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByBfqID  获取多个内容
func (obj *_BfqAnswerMgr) FetchIndexByBfqID(bfqID int) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`bfq_id` = ?", bfqID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByExamSessionID  获取多个内容
func (obj *_BfqAnswerMgr) FetchIndexByExamSessionID(examSessionID int) (results []*BfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(BfqAnswer{}).Where("`exam_session_id` = ?", examSessionID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("bfq").Where("id = ?", results[i].BfqID).Find(&results[i].Bfq).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("exam_session").Where("id = ?", results[i].ExamSessionID).Find(&results[i].ExamSession).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
