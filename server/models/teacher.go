package models

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TeacherMgr struct {
	*_BaseMgr
}

// TeacherMgr open func
func TeacherMgr(db *gorm.DB) *_TeacherMgr {
	if db == nil {
		panic(fmt.Errorf("TeacherMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TeacherMgr{_BaseMgr: &_BaseMgr{DB: db.Table("teacher"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TeacherMgr) GetTableName() string {
	return "teacher"
}

// Reset 重置gorm会话
func (obj *_TeacherMgr) Reset() *_TeacherMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TeacherMgr) Get() (result Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Teacher{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TeacherMgr) Gets() (results []*Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Teacher{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TeacherMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Teacher{}).Count(count)
}

// ////////////////////////////////////////////////////////////////////////////////

// ////////////////////////option case ////////////////////////////////////////////

// WithID id获取 用作主键
func (obj *_TeacherMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTeacherID teacher_id获取 工号
func (obj *_TeacherMgr) WithTeacherID(teacherID int16) Option {
	return optionFunc(func(o *options) { o.query["teacher_id"] = teacherID })
}

// WithName name获取 姓名
func (obj *_TeacherMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithPassword password获取 加盐后的密码
func (obj *_TeacherMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithIsAdmin is_admin获取 是否为超级管理员，0：否，1：是
func (obj *_TeacherMgr) WithIsAdmin(isAdmin []uint8) Option {
	return optionFunc(func(o *options) { o.query["is_admin"] = isAdmin })
}

// GetByOption 功能选项模式获取
func (obj *_TeacherMgr) GetByOption(opts ...Option) (result Teacher, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Teacher{}).Where(options.query).Find(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TeacherMgr) GetByOptions(opts ...Option) (results []*Teacher, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Teacher{}).Where(options.query).Find(&results).Error

	return
}

// ////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容 用作主键
func (obj *_TeacherMgr) GetFromID(id int) (result Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Teacher{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// GetBatchFromID 批量查找 用作主键
func (obj *_TeacherMgr) GetBatchFromID(ids []int) (results []*Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Teacher{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTeacherID 通过teacher_id获取内容 工号
func (obj *_TeacherMgr) GetFromTeacherID(teacherID int16) (result Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Teacher{}).Where("`teacher_id` = ?", teacherID).Find(&result).Error

	return
}

// GetBatchFromTeacherID 批量查找 工号
func (obj *_TeacherMgr) GetBatchFromTeacherID(teacherIDs []int16) (results []*Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Teacher{}).Where("`teacher_id` IN (?)", teacherIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 姓名
func (obj *_TeacherMgr) GetFromName(name string) (results []*Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Teacher{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 姓名
func (obj *_TeacherMgr) GetBatchFromName(names []string) (results []*Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Teacher{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 加盐后的密码
func (obj *_TeacherMgr) GetFromPassword(password string) (results []*Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Teacher{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 加盐后的密码
func (obj *_TeacherMgr) GetBatchFromPassword(passwords []string) (results []*Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Teacher{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromIsAdmin 通过is_admin获取内容 是否为超级管理员，0：否，1：是
func (obj *_TeacherMgr) GetFromIsAdmin(isAdmin []uint8) (results []*Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Teacher{}).Where("`is_admin` = ?", isAdmin).Find(&results).Error

	return
}

// GetBatchFromIsAdmin 批量查找 是否为超级管理员，0：否，1：是
func (obj *_TeacherMgr) GetBatchFromIsAdmin(isAdmins [][]uint8) (results []*Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Teacher{}).Where("`is_admin` IN (?)", isAdmins).Find(&results).Error

	return
}

// ////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TeacherMgr) FetchByPrimaryKey(id int) (result Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Teacher{}).Where("`id` = ?", id).Find(&result).Error

	return
}

// FetchUniqueByTeacherID primary or index 获取唯一内容
func (obj *_TeacherMgr) FetchUniqueByTeacherID(teacherID int16) (result Teacher, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Teacher{}).Where("`teacher_id` = ?", teacherID).Find(&result).Error

	return
}
