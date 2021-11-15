package dao

import (
	"github.com/gonearewe/EasyTesting/models"
)

func GetTeachersBy(teacherId string, name string, pageSize int, pageIndex int) (res []*models.Teacher) {
	var filtered = db
	if teacherId != "" {
		filtered = filtered.Where("teacher_id LIKE ?", "%"+teacherId+"%")
	}
	if name != "" {
		filtered = filtered.Where("name LIKE ?", "%"+name+"%")
	}
	filtered.Select("id", "teacher_id", "name", "is_admin").
		Limit(pageSize).Offset(pageSize * (pageIndex - 1)).Find(&res)
	return
}

func GetTeacherByTeacherId(teacherId string) *models.Teacher {
	var ret models.Teacher
	ret.TeacherID = teacherId
	db.Find(&ret)
	return &ret
}

func GetTeachersByTeacherId(pageSize int, pageIndex int) (res []models.Teacher) {
	db.Limit(pageSize).Offset(pageSize * pageIndex).Order("teacher_id").Find(&res)
	return
}

func CreateTeachers(teachers []*models.Teacher) {
	db.Create(&teachers)
}
