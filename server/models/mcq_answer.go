package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _McqAnswerMgr struct {
	*_BaseMgr
}

// McqAnswerMgr open func
func McqAnswerMgr(db *gorm.DB) *_McqAnswerMgr {
	if db == nil {
		panic(fmt.Errorf("McqAnswerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_McqAnswerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mcq_answer"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_McqAnswerMgr) GetTableName() string {
	return "mcq_answer"
}

// Reset 重置gorm会话
func (obj *_McqAnswerMgr) Reset() *_McqAnswerMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_McqAnswerMgr) Get() (result McqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("mcq").Where("id = ?", result.McqID).Find(&result.Mcq).Error; err != nil { //
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
func (obj *_McqAnswerMgr) Gets() (results []*McqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("mcq").Where("id = ?", results[i].McqID).Find(&results[i].Mcq).Error; err != nil { //
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
func (obj *_McqAnswerMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Count(count)
}

// ////////////////////////////////////////////////////////////////////////////////

// ////////////////////////option case ////////////////////////////////////////////

// WithID id获取 用作主键
func (obj *_McqAnswerMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMcqID mcq_id获取 连接 mcq
func (obj *_McqAnswerMgr) WithMcqID(mcqID int) Option {
	return optionFunc(func(o *options) { o.query["mcq_id"] = mcqID })
}

// WithExamSessionID exam_session_id获取 连接 exam_session
func (obj *_McqAnswerMgr) WithExamSessionID(examSessionID int) Option {
	return optionFunc(func(o *options) { o.query["exam_session_id"] = examSessionID })
}

// WithRightAnswer right_answer获取 正确答案，与 mcq 中同名字段保持一致
func (obj *_McqAnswerMgr) WithRightAnswer(rightAnswer string) Option {
	return optionFunc(func(o *options) { o.query["right_answer"] = rightAnswer })
}

// WithStudentAnswer student_answer获取 学生的答案
func (obj *_McqAnswerMgr) WithStudentAnswer(studentAnswer string) Option {
	return optionFunc(func(o *options) { o.query["student_answer"] = studentAnswer })
}

// WithScore score获取 本题分值
func (obj *_McqAnswerMgr) WithScore(score uint8) Option {
	return optionFunc(func(o *options) { o.query["score"] = score })
}

// GetByOption 功能选项模式获取
func (obj *_McqAnswerMgr) GetByOption(opts ...Option) (result McqAnswer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("mcq").Where("id = ?", result.McqID).Find(&result.Mcq).Error; err != nil { //
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
func (obj *_McqAnswerMgr) GetByOptions(opts ...Option) (results []*McqAnswer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("mcq").Where("id = ?", results[i].McqID).Find(&results[i].Mcq).Error; err != nil { //
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
func (obj *_McqAnswerMgr) GetFromID(id int) (result McqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("mcq").Where("id = ?", result.McqID).Find(&result.Mcq).Error; err != nil { //
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
func (obj *_McqAnswerMgr) GetBatchFromID(ids []int) (results []*McqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("mcq").Where("id = ?", results[i].McqID).Find(&results[i].Mcq).Error; err != nil { //
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

// GetFromMcqID 通过mcq_id获取内容 连接 mcq
func (obj *_McqAnswerMgr) GetFromMcqID(mcqID int) (results []*McqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Where("`mcq_id` = ?", mcqID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("mcq").Where("id = ?", results[i].McqID).Find(&results[i].Mcq).Error; err != nil { //
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

// GetBatchFromMcqID 批量查找 连接 mcq
func (obj *_McqAnswerMgr) GetBatchFromMcqID(mcqIDs []int) (results []*McqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Where("`mcq_id` IN (?)", mcqIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("mcq").Where("id = ?", results[i].McqID).Find(&results[i].Mcq).Error; err != nil { //
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
func (obj *_McqAnswerMgr) GetFromExamSessionID(examSessionID int) (results []*McqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Where("`exam_session_id` = ?", examSessionID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("mcq").Where("id = ?", results[i].McqID).Find(&results[i].Mcq).Error; err != nil { //
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
func (obj *_McqAnswerMgr) GetBatchFromExamSessionID(examSessionIDs []int) (results []*McqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Where("`exam_session_id` IN (?)", examSessionIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("mcq").Where("id = ?", results[i].McqID).Find(&results[i].Mcq).Error; err != nil { //
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

// GetFromRightAnswer 通过right_answer获取内容 正确答案，与 mcq 中同名字段保持一致
func (obj *_McqAnswerMgr) GetFromRightAnswer(rightAnswer string) (results []*McqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Where("`right_answer` = ?", rightAnswer).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("mcq").Where("id = ?", results[i].McqID).Find(&results[i].Mcq).Error; err != nil { //
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

// GetBatchFromRightAnswer 批量查找 正确答案，与 mcq 中同名字段保持一致
func (obj *_McqAnswerMgr) GetBatchFromRightAnswer(rightAnswers []string) (results []*McqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Where("`right_answer` IN (?)", rightAnswers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("mcq").Where("id = ?", results[i].McqID).Find(&results[i].Mcq).Error; err != nil { //
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

// GetFromStudentAnswer 通过student_answer获取内容 学生的答案
func (obj *_McqAnswerMgr) GetFromStudentAnswer(studentAnswer string) (results []*McqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Where("`student_answer` = ?", studentAnswer).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("mcq").Where("id = ?", results[i].McqID).Find(&results[i].Mcq).Error; err != nil { //
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

// GetBatchFromStudentAnswer 批量查找 学生的答案
func (obj *_McqAnswerMgr) GetBatchFromStudentAnswer(studentAnswers []string) (results []*McqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Where("`student_answer` IN (?)", studentAnswers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("mcq").Where("id = ?", results[i].McqID).Find(&results[i].Mcq).Error; err != nil { //
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
func (obj *_McqAnswerMgr) GetFromScore(score uint8) (results []*McqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Where("`score` = ?", score).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("mcq").Where("id = ?", results[i].McqID).Find(&results[i].Mcq).Error; err != nil { //
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
func (obj *_McqAnswerMgr) GetBatchFromScore(scores []uint8) (results []*McqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Where("`score` IN (?)", scores).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("mcq").Where("id = ?", results[i].McqID).Find(&results[i].Mcq).Error; err != nil { //
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
func (obj *_McqAnswerMgr) FetchByPrimaryKey(id int) (result McqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("mcq").Where("id = ?", result.McqID).Find(&result.Mcq).Error; err != nil { //
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

// FetchIndexByMcqID  获取多个内容
func (obj *_McqAnswerMgr) FetchIndexByMcqID(mcqID int) (results []*McqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Where("`mcq_id` = ?", mcqID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("mcq").Where("id = ?", results[i].McqID).Find(&results[i].Mcq).Error; err != nil { //
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
func (obj *_McqAnswerMgr) FetchIndexByExamSessionID(examSessionID int) (results []*McqAnswer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(McqAnswer{}).Where("`exam_session_id` = ?", examSessionID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("mcq").Where("id = ?", results[i].McqID).Find(&results[i].Mcq).Error; err != nil { //
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
