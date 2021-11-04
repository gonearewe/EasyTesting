package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _McqMgr struct {
	*_BaseMgr
}

// McqMgr open func
func McqMgr(db *gorm.DB) *_McqMgr {
	if db == nil {
		panic(fmt.Errorf("McqMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_McqMgr{_BaseMgr: &_BaseMgr{DB: db.Table("mcq"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_McqMgr) GetTableName() string {
	return "mcq"
}

// Reset 重置gorm会话
func (obj *_McqMgr) Reset() *_McqMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_McqMgr) Get() (result Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Find(&result).Error
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
func (obj *_McqMgr) Gets() (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Find(&results).Error
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
func (obj *_McqMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Mcq{}).Count(count)
}

// ////////////////////////////////////////////////////////////////////////////////

// ////////////////////////option case ////////////////////////////////////////////

// WithID id获取 用作主键
func (obj *_McqMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPublisherTeacherID publisher_teacher_id获取 创建本题的教师的工号
func (obj *_McqMgr) WithPublisherTeacherID(publisherTeacherID int16) Option {
	return optionFunc(func(o *options) { o.query["publisher_teacher_id"] = publisherTeacherID })
}

// WithStem stem获取 题干
func (obj *_McqMgr) WithStem(stem string) Option {
	return optionFunc(func(o *options) { o.query["stem"] = stem })
}

// WithChoice1 choice_1获取 选项的内容
func (obj *_McqMgr) WithChoice1(choice1 string) Option {
	return optionFunc(func(o *options) { o.query["choice_1"] = choice1 })
}

// WithChoice2 choice_2获取 选项的内容
func (obj *_McqMgr) WithChoice2(choice2 string) Option {
	return optionFunc(func(o *options) { o.query["choice_2"] = choice2 })
}

// WithChoice3 choice_3获取 选项的内容
func (obj *_McqMgr) WithChoice3(choice3 string) Option {
	return optionFunc(func(o *options) { o.query["choice_3"] = choice3 })
}

// WithChoice4 choice_4获取 选项的内容
func (obj *_McqMgr) WithChoice4(choice4 string) Option {
	return optionFunc(func(o *options) { o.query["choice_4"] = choice4 })
}

// WithChoice5 choice_5获取 选项的内容
func (obj *_McqMgr) WithChoice5(choice5 string) Option {
	return optionFunc(func(o *options) { o.query["choice_5"] = choice5 })
}

// WithChoice6 choice_6获取 选项的内容
func (obj *_McqMgr) WithChoice6(choice6 string) Option {
	return optionFunc(func(o *options) { o.query["choice_6"] = choice6 })
}

// WithChoice7 choice_7获取 选项的内容
func (obj *_McqMgr) WithChoice7(choice7 string) Option {
	return optionFunc(func(o *options) { o.query["choice_7"] = choice7 })
}

// WithIsMaq is_maq获取 是不是多选题，0：否，1：是
func (obj *_McqMgr) WithIsMaq(isMaq []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_maq"] = isMaq })
}

// WithRightAnswer right_answer获取 答案，按升序包含所有正确选项的索引，如 "5"、"124"、"67"
func (obj *_McqMgr) WithRightAnswer(rightAnswer string) Option {
	return optionFunc(func(o *options) { o.query["right_answer"] = rightAnswer })
}

// GetByOption 功能选项模式获取
func (obj *_McqMgr) GetByOption(opts ...Option) (result Mcq, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where(options.query).Find(&result).Error
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
func (obj *_McqMgr) GetByOptions(opts ...Option) (results []*Mcq, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where(options.query).Find(&results).Error
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
func (obj *_McqMgr) GetFromID(id int) (result Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`id` = ?", id).Find(&result).Error
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
func (obj *_McqMgr) GetBatchFromID(ids []int) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`id` IN (?)", ids).Find(&results).Error
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
func (obj *_McqMgr) GetFromPublisherTeacherID(publisherTeacherID int16) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`publisher_teacher_id` = ?", publisherTeacherID).Find(&results).Error
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
func (obj *_McqMgr) GetBatchFromPublisherTeacherID(publisherTeacherIDs []int16) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`publisher_teacher_id` IN (?)", publisherTeacherIDs).Find(&results).Error
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
func (obj *_McqMgr) GetFromStem(stem string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`stem` = ?", stem).Find(&results).Error
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
func (obj *_McqMgr) GetBatchFromStem(stems []string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`stem` IN (?)", stems).Find(&results).Error
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

// GetFromChoice1 通过choice_1获取内容 选项的内容
func (obj *_McqMgr) GetFromChoice1(choice1 string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`choice_1` = ?", choice1).Find(&results).Error
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

// GetBatchFromChoice1 批量查找 选项的内容
func (obj *_McqMgr) GetBatchFromChoice1(choice1s []string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`choice_1` IN (?)", choice1s).Find(&results).Error
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

// GetFromChoice2 通过choice_2获取内容 选项的内容
func (obj *_McqMgr) GetFromChoice2(choice2 string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`choice_2` = ?", choice2).Find(&results).Error
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

// GetBatchFromChoice2 批量查找 选项的内容
func (obj *_McqMgr) GetBatchFromChoice2(choice2s []string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`choice_2` IN (?)", choice2s).Find(&results).Error
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

// GetFromChoice3 通过choice_3获取内容 选项的内容
func (obj *_McqMgr) GetFromChoice3(choice3 string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`choice_3` = ?", choice3).Find(&results).Error
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

// GetBatchFromChoice3 批量查找 选项的内容
func (obj *_McqMgr) GetBatchFromChoice3(choice3s []string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`choice_3` IN (?)", choice3s).Find(&results).Error
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

// GetFromChoice4 通过choice_4获取内容 选项的内容
func (obj *_McqMgr) GetFromChoice4(choice4 string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`choice_4` = ?", choice4).Find(&results).Error
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

// GetBatchFromChoice4 批量查找 选项的内容
func (obj *_McqMgr) GetBatchFromChoice4(choice4s []string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`choice_4` IN (?)", choice4s).Find(&results).Error
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

// GetFromChoice5 通过choice_5获取内容 选项的内容
func (obj *_McqMgr) GetFromChoice5(choice5 string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`choice_5` = ?", choice5).Find(&results).Error
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

// GetBatchFromChoice5 批量查找 选项的内容
func (obj *_McqMgr) GetBatchFromChoice5(choice5s []string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`choice_5` IN (?)", choice5s).Find(&results).Error
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

// GetFromChoice6 通过choice_6获取内容 选项的内容
func (obj *_McqMgr) GetFromChoice6(choice6 string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`choice_6` = ?", choice6).Find(&results).Error
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

// GetBatchFromChoice6 批量查找 选项的内容
func (obj *_McqMgr) GetBatchFromChoice6(choice6s []string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`choice_6` IN (?)", choice6s).Find(&results).Error
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

// GetFromChoice7 通过choice_7获取内容 选项的内容
func (obj *_McqMgr) GetFromChoice7(choice7 string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`choice_7` = ?", choice7).Find(&results).Error
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

// GetBatchFromChoice7 批量查找 选项的内容
func (obj *_McqMgr) GetBatchFromChoice7(choice7s []string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`choice_7` IN (?)", choice7s).Find(&results).Error
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

// GetFromIsMaq 通过is_maq获取内容 是不是多选题，0：否，1：是
func (obj *_McqMgr) GetFromIsMaq(isMaq []uint8) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`is_maq` = ?", isMaq).Find(&results).Error
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

// GetBatchFromIsMaq 批量查找 是不是多选题，0：否，1：是
func (obj *_McqMgr) GetBatchFromIsMaq(isMaqs [][]uint8) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`is_maq` IN (?)", isMaqs).Find(&results).Error
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

// GetFromRightAnswer 通过right_answer获取内容 答案，按升序包含所有正确选项的索引，如 "5"、"124"、"67"
func (obj *_McqMgr) GetFromRightAnswer(rightAnswer string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`right_answer` = ?", rightAnswer).Find(&results).Error
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

// GetBatchFromRightAnswer 批量查找 答案，按升序包含所有正确选项的索引，如 "5"、"124"、"67"
func (obj *_McqMgr) GetBatchFromRightAnswer(rightAnswers []string) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`right_answer` IN (?)", rightAnswers).Find(&results).Error
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
func (obj *_McqMgr) FetchByPrimaryKey(id int) (result Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`id` = ?", id).Find(&result).Error
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
func (obj *_McqMgr) FetchIndexByPublisherTeacherID(publisherTeacherID int16) (results []*Mcq, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Mcq{}).Where("`publisher_teacher_id` = ?", publisherTeacherID).Find(&results).Error
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
