package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TfqAnswerMgr struct {
	*_BaseMgr
}

// TfqAnswerMgr open func
func TfqAnswerMgr(db *gorm.DB) *_TfqAnswerMgr {
	if db == nil {
		panic(fmt.Errorf("TfqAnswerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TfqAnswerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tfq_answer"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TfqAnswerMgr) GetTableName() string {
	return "tfq_answer"
}

// Reset 重置gorm会话
func (obj *_TfqAnswerMgr) Reset() *_TfqAnswerMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TfqAnswerMgr) Get() (result TfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("tfq").Where("id = ?", result.TfqID).Find(&result.Tfq).Error; err != nil { //
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
func (obj *_TfqAnswerMgr) Gets() (results []*TfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tfq").Where("id = ?", results[i].TfqID).Find(&results[i].Tfq).Error; err != nil { //
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
func (obj *_TfqAnswerMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Count(count)
}

// ////////////////////////////////////////////////////////////////////////////////

// ////////////////////////option case ////////////////////////////////////////////

// WithID id获取 用作主键
func (obj *_TfqAnswerMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTfqID tfq_id获取 连接 tfq
func (obj *_TfqAnswerMgr) WithTfqID(tfqID int) Option {
	return optionFunc(func(o *options) { o.query["tfq_id"] = tfqID })
}

// WithExamSessionID exam_session_id获取 连接 exam_session
func (obj *_TfqAnswerMgr) WithExamSessionID(examSessionID int) Option {
	return optionFunc(func(o *options) { o.query["exam_session_id"] = examSessionID })
}

// WithRightAnswer right_answer获取 正确答案，与 tfq 中同名字段保持一致
func (obj *_TfqAnswerMgr) WithRightAnswer(rightAnswer []uint8) Option {
	return optionFunc(func(o *options) { o.query["right_answer"] = rightAnswer })
}

// WithStudentAnswer student_answer获取 学生的答案，0：错，1：对
func (obj *_TfqAnswerMgr) WithStudentAnswer(studentAnswer []uint8) Option {
	return optionFunc(func(o *options) { o.query["student_answer"] = studentAnswer })
}

// WithScore score获取 本题分值
func (obj *_TfqAnswerMgr) WithScore(score uint8) Option {
	return optionFunc(func(o *options) { o.query["score"] = score })
}

// GetByOption 功能选项模式获取
func (obj *_TfqAnswerMgr) GetByOption(opts ...Option) (result TfqAnswer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("tfq").Where("id = ?", result.TfqID).Find(&result.Tfq).Error; err != nil { //
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
func (obj *_TfqAnswerMgr) GetByOptions(opts ...Option) (results []*TfqAnswer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tfq").Where("id = ?", results[i].TfqID).Find(&results[i].Tfq).Error; err != nil { //
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
func (obj *_TfqAnswerMgr) GetFromID(id int) (result TfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("tfq").Where("id = ?", result.TfqID).Find(&result.Tfq).Error; err != nil { //
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
func (obj *_TfqAnswerMgr) GetBatchFromID(ids []int) (results []*TfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tfq").Where("id = ?", results[i].TfqID).Find(&results[i].Tfq).Error; err != nil { //
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

// GetFromTfqID 通过tfq_id获取内容 连接 tfq
func (obj *_TfqAnswerMgr) GetFromTfqID(tfqID int) (results []*TfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Where("`tfq_id` = ?", tfqID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tfq").Where("id = ?", results[i].TfqID).Find(&results[i].Tfq).Error; err != nil { //
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

// GetBatchFromTfqID 批量查找 连接 tfq
func (obj *_TfqAnswerMgr) GetBatchFromTfqID(tfqIDs []int) (results []*TfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Where("`tfq_id` IN (?)", tfqIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tfq").Where("id = ?", results[i].TfqID).Find(&results[i].Tfq).Error; err != nil { //
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
func (obj *_TfqAnswerMgr) GetFromExamSessionID(examSessionID int) (results []*TfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Where("`exam_session_id` = ?", examSessionID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tfq").Where("id = ?", results[i].TfqID).Find(&results[i].Tfq).Error; err != nil { //
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
func (obj *_TfqAnswerMgr) GetBatchFromExamSessionID(examSessionIDs []int) (results []*TfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Where("`exam_session_id` IN (?)", examSessionIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tfq").Where("id = ?", results[i].TfqID).Find(&results[i].Tfq).Error; err != nil { //
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

// GetFromRightAnswer 通过right_answer获取内容 正确答案，与 tfq 中同名字段保持一致
func (obj *_TfqAnswerMgr) GetFromRightAnswer(rightAnswer []uint8) (results []*TfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Where("`right_answer` = ?", rightAnswer).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tfq").Where("id = ?", results[i].TfqID).Find(&results[i].Tfq).Error; err != nil { //
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

// GetBatchFromRightAnswer 批量查找 正确答案，与 tfq 中同名字段保持一致
func (obj *_TfqAnswerMgr) GetBatchFromRightAnswer(rightAnswers [][]uint8) (results []*TfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Where("`right_answer` IN (?)", rightAnswers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tfq").Where("id = ?", results[i].TfqID).Find(&results[i].Tfq).Error; err != nil { //
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

// GetFromStudentAnswer 通过student_answer获取内容 学生的答案，0：错，1：对
func (obj *_TfqAnswerMgr) GetFromStudentAnswer(studentAnswer []uint8) (results []*TfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Where("`student_answer` = ?", studentAnswer).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tfq").Where("id = ?", results[i].TfqID).Find(&results[i].Tfq).Error; err != nil { //
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

// GetBatchFromStudentAnswer 批量查找 学生的答案，0：错，1：对
func (obj *_TfqAnswerMgr) GetBatchFromStudentAnswer(studentAnswers [][]uint8) (results []*TfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Where("`student_answer` IN (?)", studentAnswers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tfq").Where("id = ?", results[i].TfqID).Find(&results[i].Tfq).Error; err != nil { //
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
func (obj *_TfqAnswerMgr) GetFromScore(score uint8) (results []*TfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Where("`score` = ?", score).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tfq").Where("id = ?", results[i].TfqID).Find(&results[i].Tfq).Error; err != nil { //
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
func (obj *_TfqAnswerMgr) GetBatchFromScore(scores []uint8) (results []*TfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Where("`score` IN (?)", scores).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tfq").Where("id = ?", results[i].TfqID).Find(&results[i].Tfq).Error; err != nil { //
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
func (obj *_TfqAnswerMgr) FetchByPrimaryKey(id int) (result TfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("tfq").Where("id = ?", result.TfqID).Find(&result.Tfq).Error; err != nil { //
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

// FetchIndexByTfqID  获取多个内容
func (obj *_TfqAnswerMgr) FetchIndexByTfqID(tfqID int) (results []*TfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Where("`tfq_id` = ?", tfqID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tfq").Where("id = ?", results[i].TfqID).Find(&results[i].Tfq).Error; err != nil { //
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
func (obj *_TfqAnswerMgr) FetchIndexByExamSessionID(examSessionID int) (results []*TfqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TfqAnswer{}).Where("`exam_session_id` = ?", examSessionID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("tfq").Where("id = ?", results[i].TfqID).Find(&results[i].Tfq).Error; err != nil { //
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
