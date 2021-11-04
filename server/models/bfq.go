package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _BfqMgr struct {
	*_BaseMgr
}

// BfqMgr open func
func BfqMgr(db *gorm.DB) *_BfqMgr {
	if db == nil {
		panic(fmt.Errorf("BfqMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_BfqMgr{_BaseMgr: &_BaseMgr{DB: db.Table("bfq"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_BfqMgr) GetTableName() string {
	return "bfq"
}

// Reset 重置gorm会话
func (obj *_BfqMgr) Reset() *_BfqMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_BfqMgr) Get() (result Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Find(&result).Error
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
func (obj *_BfqMgr) Gets() (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Find(&results).Error
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
func (obj *_BfqMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Bfq{}).Count(count)
}

// ////////////////////////////////////////////////////////////////////////////////

// ////////////////////////option case ////////////////////////////////////////////

// WithID id获取 用作主键
func (obj *_BfqMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPublisherTeacherID publisher_teacher_id获取 创建本题的教师的工号
func (obj *_BfqMgr) WithPublisherTeacherID(publisherTeacherID int16) Option {
	return optionFunc(func(o *options) { o.query["publisher_teacher_id"] = publisherTeacherID })
}

// WithStem stem获取 题干
func (obj *_BfqMgr) WithStem(stem string) Option {
	return optionFunc(func(o *options) { o.query["stem"] = stem })
}

// WithBlankNum blank_num获取 要填的空的数目，若大于 1，则说明是 crq
func (obj *_BfqMgr) WithBlankNum(blankNum bool) Option {
	return optionFunc(func(o *options) { o.query["blank_num"] = blankNum })
}

// WithAnswer1 answer_1获取 填空的答案
func (obj *_BfqMgr) WithAnswer1(answer1 string) Option {
	return optionFunc(func(o *options) { o.query["answer_1"] = answer1 })
}

// WithAnswer2 answer_2获取 填空的答案
func (obj *_BfqMgr) WithAnswer2(answer2 string) Option {
	return optionFunc(func(o *options) { o.query["answer_2"] = answer2 })
}

// WithAnswer3 answer_3获取 填空的答案
func (obj *_BfqMgr) WithAnswer3(answer3 string) Option {
	return optionFunc(func(o *options) { o.query["answer_3"] = answer3 })
}

// WithAnswer4 answer_4获取 填空的答案
func (obj *_BfqMgr) WithAnswer4(answer4 string) Option {
	return optionFunc(func(o *options) { o.query["answer_4"] = answer4 })
}

// WithAnswer5 answer_5获取 填空的答案
func (obj *_BfqMgr) WithAnswer5(answer5 string) Option {
	return optionFunc(func(o *options) { o.query["answer_5"] = answer5 })
}

// GetByOption 功能选项模式获取
func (obj *_BfqMgr) GetByOption(opts ...Option) (result Bfq, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where(options.query).Find(&result).Error
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
func (obj *_BfqMgr) GetByOptions(opts ...Option) (results []*Bfq, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where(options.query).Find(&results).Error
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
func (obj *_BfqMgr) GetFromID(id int) (result Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`id` = ?", id).Find(&result).Error
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
func (obj *_BfqMgr) GetBatchFromID(ids []int) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`id` IN (?)", ids).Find(&results).Error
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
func (obj *_BfqMgr) GetFromPublisherTeacherID(publisherTeacherID int16) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`publisher_teacher_id` = ?", publisherTeacherID).Find(&results).Error
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
func (obj *_BfqMgr) GetBatchFromPublisherTeacherID(publisherTeacherIDs []int16) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`publisher_teacher_id` IN (?)", publisherTeacherIDs).Find(&results).Error
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
func (obj *_BfqMgr) GetFromStem(stem string) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`stem` = ?", stem).Find(&results).Error
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
func (obj *_BfqMgr) GetBatchFromStem(stems []string) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`stem` IN (?)", stems).Find(&results).Error
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

// GetFromBlankNum 通过blank_num获取内容 要填的空的数目，若大于 1，则说明是 crq
func (obj *_BfqMgr) GetFromBlankNum(blankNum bool) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`blank_num` = ?", blankNum).Find(&results).Error
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

// GetBatchFromBlankNum 批量查找 要填的空的数目，若大于 1，则说明是 crq
func (obj *_BfqMgr) GetBatchFromBlankNum(blankNums []bool) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`blank_num` IN (?)", blankNums).Find(&results).Error
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

// GetFromAnswer1 通过answer_1获取内容 填空的答案
func (obj *_BfqMgr) GetFromAnswer1(answer1 string) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`answer_1` = ?", answer1).Find(&results).Error
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

// GetBatchFromAnswer1 批量查找 填空的答案
func (obj *_BfqMgr) GetBatchFromAnswer1(answer1s []string) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`answer_1` IN (?)", answer1s).Find(&results).Error
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

// GetFromAnswer2 通过answer_2获取内容 填空的答案
func (obj *_BfqMgr) GetFromAnswer2(answer2 string) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`answer_2` = ?", answer2).Find(&results).Error
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

// GetBatchFromAnswer2 批量查找 填空的答案
func (obj *_BfqMgr) GetBatchFromAnswer2(answer2s []string) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`answer_2` IN (?)", answer2s).Find(&results).Error
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

// GetFromAnswer3 通过answer_3获取内容 填空的答案
func (obj *_BfqMgr) GetFromAnswer3(answer3 string) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`answer_3` = ?", answer3).Find(&results).Error
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

// GetBatchFromAnswer3 批量查找 填空的答案
func (obj *_BfqMgr) GetBatchFromAnswer3(answer3s []string) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`answer_3` IN (?)", answer3s).Find(&results).Error
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

// GetFromAnswer4 通过answer_4获取内容 填空的答案
func (obj *_BfqMgr) GetFromAnswer4(answer4 string) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`answer_4` = ?", answer4).Find(&results).Error
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

// GetBatchFromAnswer4 批量查找 填空的答案
func (obj *_BfqMgr) GetBatchFromAnswer4(answer4s []string) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`answer_4` IN (?)", answer4s).Find(&results).Error
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

// GetFromAnswer5 通过answer_5获取内容 填空的答案
func (obj *_BfqMgr) GetFromAnswer5(answer5 string) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`answer_5` = ?", answer5).Find(&results).Error
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

// GetBatchFromAnswer5 批量查找 填空的答案
func (obj *_BfqMgr) GetBatchFromAnswer5(answer5s []string) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`answer_5` IN (?)", answer5s).Find(&results).Error
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
func (obj *_BfqMgr) FetchByPrimaryKey(id int) (result Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`id` = ?", id).Find(&result).Error
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
func (obj *_BfqMgr) FetchIndexByPublisherTeacherID(publisherTeacherID int16) (results []*Bfq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Bfq{}).Where("`publisher_teacher_id` = ?", publisherTeacherID).Find(&results).Error
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
