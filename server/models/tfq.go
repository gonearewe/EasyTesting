package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TfqMgr struct {
	*_BaseMgr
}

// TfqMgr open func
func TfqMgr(db *gorm.DB) *_TfqMgr {
	if db == nil {
		panic(fmt.Errorf("TfqMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TfqMgr{_BaseMgr: &_BaseMgr{DB: db.Table("tfq"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TfqMgr) GetTableName() string {
	return "tfq"
}

// Reset 重置gorm会话
func (obj *_TfqMgr) Reset() *_TfqMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TfqMgr) Get() (result Tfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tfq{}).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("teacher").Where("teacher_id = ?", result.PublisherTeacherID).Find(&result.Teacher).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_TfqMgr) Gets() (results []*Tfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tfq{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("teacher").Where("teacher_id = ?", results[i].PublisherTeacherID).Find(&results[i].Teacher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TfqMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Tfq{}).Count(count)
}

// ////////////////////////////////////////////////////////////////////////////////

// ////////////////////////option case ////////////////////////////////////////////

// WithID id获取 用作主键
func (obj *_TfqMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPublisherTeacherID publisher_teacher_id获取 创建本题的教师的工号
func (obj *_TfqMgr) WithPublisherTeacherID(publisherTeacherID int16) Option {
	return optionFunc(func(o *options) { o.query["publisher_teacher_id"] = publisherTeacherID })
}

// WithStem stem获取 题干
func (obj *_TfqMgr) WithStem(stem string) Option {
	return optionFunc(func(o *options) { o.query["stem"] = stem })
}

// WithAnswer answer获取 0：错，1：对
func (obj *_TfqMgr) WithAnswer(answer []uint8) Option {
	return optionFunc(func(o *options) { o.query["answer"] = answer })
}

// GetByOption 功能选项模式获取
func (obj *_TfqMgr) GetByOption(opts ...Option) (result Tfq, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Tfq{}).Where(options.query).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("teacher").Where("teacher_id = ?", result.PublisherTeacherID).Find(&result.Teacher).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TfqMgr) GetByOptions(opts ...Option) (results []*Tfq, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Tfq{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("teacher").Where("teacher_id = ?", results[i].PublisherTeacherID).Find(&results[i].Teacher).Error; err != nil { //
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
func (obj *_TfqMgr) GetFromID(id int) (result Tfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tfq{}).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("teacher").Where("teacher_id = ?", result.PublisherTeacherID).Find(&result.Teacher).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromID 批量查找 用作主键
func (obj *_TfqMgr) GetBatchFromID(ids []int) (results []*Tfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tfq{}).Where("`id` IN (?)", ids).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("teacher").Where("teacher_id = ?", results[i].PublisherTeacherID).Find(&results[i].Teacher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromPublisherTeacherID 通过publisher_teacher_id获取内容 创建本题的教师的工号
func (obj *_TfqMgr) GetFromPublisherTeacherID(publisherTeacherID int16) (results []*Tfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tfq{}).Where("`publisher_teacher_id` = ?", publisherTeacherID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("teacher").Where("teacher_id = ?", results[i].PublisherTeacherID).Find(&results[i].Teacher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromPublisherTeacherID 批量查找 创建本题的教师的工号
func (obj *_TfqMgr) GetBatchFromPublisherTeacherID(publisherTeacherIDs []int16) (results []*Tfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tfq{}).Where("`publisher_teacher_id` IN (?)", publisherTeacherIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("teacher").Where("teacher_id = ?", results[i].PublisherTeacherID).Find(&results[i].Teacher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromStem 通过stem获取内容 题干
func (obj *_TfqMgr) GetFromStem(stem string) (results []*Tfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tfq{}).Where("`stem` = ?", stem).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("teacher").Where("teacher_id = ?", results[i].PublisherTeacherID).Find(&results[i].Teacher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromStem 批量查找 题干
func (obj *_TfqMgr) GetBatchFromStem(stems []string) (results []*Tfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tfq{}).Where("`stem` IN (?)", stems).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("teacher").Where("teacher_id = ?", results[i].PublisherTeacherID).Find(&results[i].Teacher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromAnswer 通过answer获取内容 0：错，1：对
func (obj *_TfqMgr) GetFromAnswer(answer []uint8) (results []*Tfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tfq{}).Where("`answer` = ?", answer).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("teacher").Where("teacher_id = ?", results[i].PublisherTeacherID).Find(&results[i].Teacher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromAnswer 批量查找 0：错，1：对
func (obj *_TfqMgr) GetBatchFromAnswer(answers [][]uint8) (results []*Tfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tfq{}).Where("`answer` IN (?)", answers).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("teacher").Where("teacher_id = ?", results[i].PublisherTeacherID).Find(&results[i].Teacher).Error; err != nil { //
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
func (obj *_TfqMgr) FetchByPrimaryKey(id int) (result Tfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tfq{}).Where("`id` = ?", id).Find(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("teacher").Where("teacher_id = ?", result.PublisherTeacherID).Find(&result.Teacher).Error; err != nil { //
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchIndexByPublisherTeacherID  获取多个内容
func (obj *_TfqMgr) FetchIndexByPublisherTeacherID(publisherTeacherID int16) (results []*Tfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Tfq{}).Where("`publisher_teacher_id` = ?", publisherTeacherID).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("teacher").Where("teacher_id = ?", results[i].PublisherTeacherID).Find(&results[i].Teacher).Error; err != nil { //
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}
