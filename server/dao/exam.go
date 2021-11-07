package dao

import "github.com/gonearewe/EasyTesting/models"

func GetExamsById(pageSize int, pageIndex int) (res []models.Exam, err error) {
	err = db.Limit(pageSize).Offset(pageSize * pageIndex).Find(&res).Error
	return
}

func InsertExam(exam models.Exam) error {
	return db.Create(exam).Error
}
