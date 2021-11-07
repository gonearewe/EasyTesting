package dao

import "github.com/gonearewe/EasyTesting/models"

func GetStudentsById(pageSize int, pageIndex int) (res []models.Student) {
	db.Limit(pageSize).Offset(pageSize * pageIndex).Find(&res)
	return
}

func GetStudentsByStudentId(pageSize int, pageIndex int) (res []models.Student) {
	db.Limit(pageSize).Offset(pageSize * pageIndex).Order("student_id").Find(&res)
	return
}
