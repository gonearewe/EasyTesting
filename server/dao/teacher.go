package dao

import "github.com/gonearewe/EasyTesting/models"

func GetTeachersById(pageSize int, pageIndex int) (res []models.Teacher) {
	db.Limit(pageSize).Offset(pageSize * pageIndex).Find(&res)
	return
}

func GetTeacherByTeacherId(teacherId string) (res models.Teacher) {
	res.TeacherID = teacherId
	db.Find(&res)
	return
}

func GetTeachersByTeacherId(pageSize int, pageIndex int) (res []models.Teacher) {
	db.Limit(pageSize).Offset(pageSize * pageIndex).Order("teacher_id").Find(&res)
	return
}

func CreateTeachers(teachers []*models.Teacher) {
	db.Create(&teachers)
}
