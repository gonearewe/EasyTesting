package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ExamSessionMgr struct {
	*_BaseMgr
}

// ExamSessionMgr open func
func ExamSessionMgr(db *gorm.DB) *_ExamSessionMgr {
	if db == nil {
		panic(fmt.Errorf("ExamSessionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ExamSessionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("exam_session"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ExamSessionMgr) GetTableName() string {
	return "exam_session"
}

// Reset 重置gorm会话
func (obj *_ExamSessionMgr) Reset() *_ExamSessionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ExamSessionMgr) Get() (result ExamSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("exam").Where("id = ?", result.ExamID).Find(&result.Exam).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("student").Where("student_id = ?", result.StudentID).Find(&result.Student).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_ExamSessionMgr) Gets() (results []*ExamSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("exam").Where("id = ?", results[i].ExamID).Find(&results[i].Exam).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("student").Where("student_id = ?", results[i].StudentID).Find(&results[i].Student).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ExamSessionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Count(count)
}

// ////////////////////////////////////////////////////////////////////////////////

// ////////////////////////option case ////////////////////////////////////////////

// WithID id获取 用作主键
func (obj *_ExamSessionMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithExamID exam_id获取 连接 exam
func (obj *_ExamSessionMgr) WithExamID(examID int) Option {
	return optionFunc(func(o *options) { o.query["exam_id"] = examID })
}

// WithStudentID student_id获取 连接 student
func (obj *_ExamSessionMgr) WithStudentID(studentID int16) Option {
	return optionFunc(func(o *options) { o.query["student_id"] = studentID })
}

// WithStartTime start_time获取 作答开始时间
func (obj *_ExamSessionMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取 交卷时间
func (obj *_ExamSessionMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithScore score获取 最终成绩
func (obj *_ExamSessionMgr) WithScore(score uint8) Option {
	return optionFunc(func(o *options) { o.query["score"] = score })
}

// GetByOption 功能选项模式获取
func (obj *_ExamSessionMgr) GetByOption(opts ...Option) (result ExamSession, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("exam").Where("id = ?", result.ExamID).Find(&result.Exam).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("student").Where("student_id = ?", result.StudentID).Find(&result.Student).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ExamSessionMgr) GetByOptions(opts ...Option) (results []*ExamSession, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("exam").Where("id = ?", results[i].ExamID).Find(&results[i].Exam).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("student").Where("student_id = ?", results[i].StudentID).Find(&results[i].Student).Error; err != nil { //
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
func (obj *_ExamSessionMgr) GetFromID(id int) (result ExamSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("exam").Where("id = ?", result.ExamID).Find(&result.Exam).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("student").Where("student_id = ?", result.StudentID).Find(&result.Student).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找 用作主键
func (obj *_ExamSessionMgr) GetBatchFromID(ids []int) (results []*ExamSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("exam").Where("id = ?", results[i].ExamID).Find(&results[i].Exam).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("student").Where("student_id = ?", results[i].StudentID).Find(&results[i].Student).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromExamID 通过exam_id获取内容 连接 exam
func (obj *_ExamSessionMgr) GetFromExamID(examID int) (results []*ExamSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Where("`exam_id` = ?", examID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("exam").Where("id = ?", results[i].ExamID).Find(&results[i].Exam).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("student").Where("student_id = ?", results[i].StudentID).Find(&results[i].Student).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromExamID 批量查找 连接 exam
func (obj *_ExamSessionMgr) GetBatchFromExamID(examIDs []int) (results []*ExamSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Where("`exam_id` IN (?)", examIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("exam").Where("id = ?", results[i].ExamID).Find(&results[i].Exam).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("student").Where("student_id = ?", results[i].StudentID).Find(&results[i].Student).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStudentID 通过student_id获取内容 连接 student
func (obj *_ExamSessionMgr) GetFromStudentID(studentID int16) (results []*ExamSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Where("`student_id` = ?", studentID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("exam").Where("id = ?", results[i].ExamID).Find(&results[i].Exam).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("student").Where("student_id = ?", results[i].StudentID).Find(&results[i].Student).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStudentID 批量查找 连接 student
func (obj *_ExamSessionMgr) GetBatchFromStudentID(studentIDs []int16) (results []*ExamSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Where("`student_id` IN (?)", studentIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("exam").Where("id = ?", results[i].ExamID).Find(&results[i].Exam).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("student").Where("student_id = ?", results[i].StudentID).Find(&results[i].Student).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStartTime 通过start_time获取内容 作答开始时间
func (obj *_ExamSessionMgr) GetFromStartTime(startTime time.Time) (results []*ExamSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Where("`start_time` = ?", startTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("exam").Where("id = ?", results[i].ExamID).Find(&results[i].Exam).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("student").Where("student_id = ?", results[i].StudentID).Find(&results[i].Student).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStartTime 批量查找 作答开始时间
func (obj *_ExamSessionMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*ExamSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("exam").Where("id = ?", results[i].ExamID).Find(&results[i].Exam).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("student").Where("student_id = ?", results[i].StudentID).Find(&results[i].Student).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromEndTime 通过end_time获取内容 交卷时间
func (obj *_ExamSessionMgr) GetFromEndTime(endTime time.Time) (results []*ExamSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Where("`end_time` = ?", endTime).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("exam").Where("id = ?", results[i].ExamID).Find(&results[i].Exam).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("student").Where("student_id = ?", results[i].StudentID).Find(&results[i].Student).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromEndTime 批量查找 交卷时间
func (obj *_ExamSessionMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*ExamSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("exam").Where("id = ?", results[i].ExamID).Find(&results[i].Exam).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("student").Where("student_id = ?", results[i].StudentID).Find(&results[i].Student).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromScore 通过score获取内容 最终成绩
func (obj *_ExamSessionMgr) GetFromScore(score uint8) (results []*ExamSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Where("`score` = ?", score).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("exam").Where("id = ?", results[i].ExamID).Find(&results[i].Exam).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("student").Where("student_id = ?", results[i].StudentID).Find(&results[i].Student).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromScore 批量查找 最终成绩
func (obj *_ExamSessionMgr) GetBatchFromScore(scores []uint8) (results []*ExamSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Where("`score` IN (?)", scores).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("exam").Where("id = ?", results[i].ExamID).Find(&results[i].Exam).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("student").Where("student_id = ?", results[i].StudentID).Find(&results[i].Student).Error; err != nil { //
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
func (obj *_ExamSessionMgr) FetchByPrimaryKey(id int) (result ExamSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("exam").Where("id = ?", result.ExamID).Find(&result.Exam).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
		if err = obj.NewDB().Table("student").Where("student_id = ?", result.StudentID).Find(&result.Student).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByExamID  获取多个内容
func (obj *_ExamSessionMgr) FetchIndexByExamID(examID int) (results []*ExamSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Where("`exam_id` = ?", examID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("exam").Where("id = ?", results[i].ExamID).Find(&results[i].Exam).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("student").Where("student_id = ?", results[i].StudentID).Find(&results[i].Student).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// FetchIndexByStudentID  获取多个内容
func (obj *_ExamSessionMgr) FetchIndexByStudentID(studentID int16) (results []*ExamSession, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ExamSession{}).Where("`student_id` = ?", studentID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("exam").Where("id = ?", results[i].ExamID).Find(&results[i].Exam).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
			if err = obj.NewDB().Table("student").Where("student_id = ?", results[i].StudentID).Find(&results[i].Student).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
