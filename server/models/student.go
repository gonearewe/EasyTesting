package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _StudentMgr struct {
	*_BaseMgr
}

// StudentMgr open func
func StudentMgr(db *gorm.DB) *_StudentMgr {
	if db == nil {
		panic(fmt.Errorf("StudentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_StudentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("student"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_StudentMgr) GetTableName() string {
	return "student"
}

// Reset 重置gorm会话
func (obj *_StudentMgr) Reset() *_StudentMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_StudentMgr) Get() (result Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Student{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_StudentMgr) Gets() (results []*Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Student{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_StudentMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Student{}).Count(count)
}

// ////////////////////////////////////////////////////////////////////////////////

// ////////////////////////option case ////////////////////////////////////////////

// WithID id获取 用作主键
func (obj *_StudentMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStudentID student_id获取 学号
func (obj *_StudentMgr) WithStudentID(studentID int16) Option {
	return optionFunc(func(o *options) { o.query["student_id"] = studentID })
}

// WithName name获取 姓名
func (obj *_StudentMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithClassID class_id获取 班号
func (obj *_StudentMgr) WithClassID(classID int16) Option {
	return optionFunc(func(o *options) { o.query["class_id"] = classID })
}

// GetByOption 功能选项模式获取
func (obj *_StudentMgr) GetByOption(opts ...Option) (result Student, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Student{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_StudentMgr) GetByOptions(opts ...Option) (results []*Student, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Student{}).Where(options.query).Find(&results).Error

	return
}

// ////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 用作主键
func (obj *_StudentMgr) GetFromID(id int) (result Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Student{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 用作主键
func (obj *_StudentMgr) GetBatchFromID(ids []int) (results []*Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Student{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromStudentID 通过student_id获取内容 学号
func (obj *_StudentMgr) GetFromStudentID(studentID int16) (result Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Student{}).Where("`student_id` = ?", studentID).Find(&result).Error

	return
}

// GetBatchFromStudentID 批量查找 学号
func (obj *_StudentMgr) GetBatchFromStudentID(studentIDs []int16) (results []*Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Student{}).Where("`student_id` IN (?)", studentIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 姓名
func (obj *_StudentMgr) GetFromName(name string) (results []*Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Student{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 姓名
func (obj *_StudentMgr) GetBatchFromName(names []string) (results []*Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Student{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromClassID 通过class_id获取内容 班号
func (obj *_StudentMgr) GetFromClassID(classID int16) (results []*Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Student{}).Where("`class_id` = ?", classID).Find(&results).Error

	return
}

// GetBatchFromClassID 批量查找 班号
func (obj *_StudentMgr) GetBatchFromClassID(classIDs []int16) (results []*Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Student{}).Where("`class_id` IN (?)", classIDs).Find(&results).Error

	return
}

// ////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_StudentMgr) FetchByPrimaryKey(id int) (result Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Student{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByStudentID primary or index 获取唯一内容
func (obj *_StudentMgr) FetchUniqueByStudentID(studentID int16) (result Student, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Student{}).Where("`student_id` = ?", studentID).Find(&result).Error

	return
}
