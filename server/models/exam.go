package models

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _ExamMgr struct {
	*_BaseMgr
}

// ExamMgr open func
func ExamMgr(db *gorm.DB) *_ExamMgr {
	if db == nil {
		panic(fmt.Errorf("ExamMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_ExamMgr{_BaseMgr: &_BaseMgr{DB: db.Table("exam"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ExamMgr) GetTableName() string {
	return "exam"
}

// Reset 重置gorm会话
func (obj *_ExamMgr) Reset() *_ExamMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ExamMgr) Get() (result Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Exam{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ExamMgr) Gets() (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Exam{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_ExamMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Exam{}).Count(count)
}

// ////////////////////////////////////////////////////////////////////////////////

// ////////////////////////option case ////////////////////////////////////////////

// WithID id获取 用作主键
func (obj *_ExamMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPublisherTeacherID publisher_teacher_id获取 发布考试的教师的工号
func (obj *_ExamMgr) WithPublisherTeacherID(publisherTeacherID int16) Option {
	return optionFunc(func(o *options) { o.query["publisher_teacher_id"] = publisherTeacherID })
}

// WithStartTime start_time获取 考试开始时间
func (obj *_ExamMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取 考试结束时间
func (obj *_ExamMgr) WithEndTime(endTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithTimeAllowed time_allowed获取 考生答题时间
func (obj *_ExamMgr) WithTimeAllowed(timeAllowed string) Option {
	return optionFunc(func(o *options) { o.query["time_allowed"] = timeAllowed })
}

// GetByOption 功能选项模式获取
func (obj *_ExamMgr) GetByOption(opts ...Option) (result Exam, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Exam{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_ExamMgr) GetByOptions(opts ...Option) (results []*Exam, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Exam{}).Where(options.query).Find(&results).Error

	return
}

// ////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 用作主键
func (obj *_ExamMgr) GetFromID(id int) (result Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Exam{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 用作主键
func (obj *_ExamMgr) GetBatchFromID(ids []int) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Exam{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPublisherTeacherID 通过publisher_teacher_id获取内容 发布考试的教师的工号
func (obj *_ExamMgr) GetFromPublisherTeacherID(publisherTeacherID int16) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Exam{}).Where("`publisher_teacher_id` = ?", publisherTeacherID).Find(&results).Error

	return
}

// GetBatchFromPublisherTeacherID 批量查找 发布考试的教师的工号
func (obj *_ExamMgr) GetBatchFromPublisherTeacherID(publisherTeacherIDs []int16) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Exam{}).Where("`publisher_teacher_id` IN (?)", publisherTeacherIDs).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容 考试开始时间
func (obj *_ExamMgr) GetFromStartTime(startTime time.Time) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Exam{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找 考试开始时间
func (obj *_ExamMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Exam{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 考试结束时间
func (obj *_ExamMgr) GetFromEndTime(endTime time.Time) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Exam{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 考试结束时间
func (obj *_ExamMgr) GetBatchFromEndTime(endTimes []time.Time) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Exam{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromTimeAllowed 通过time_allowed获取内容 考生答题时间
func (obj *_ExamMgr) GetFromTimeAllowed(timeAllowed string) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Exam{}).Where("`time_allowed` = ?", timeAllowed).Find(&results).Error

	return
}

// GetBatchFromTimeAllowed 批量查找 考生答题时间
func (obj *_ExamMgr) GetBatchFromTimeAllowed(timeAlloweds []string) (results []*Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Exam{}).Where("`time_allowed` IN (?)", timeAlloweds).Find(&results).Error

	return
}

// ////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_ExamMgr) FetchByPrimaryKey(id int) (result Exam, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Exam{}).Where("`id` = ?", id).Find(&result).Error

	return
}
