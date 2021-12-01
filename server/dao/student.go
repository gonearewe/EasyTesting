package dao

import (
	"github.com/gonearewe/EasyTesting/models"
	"gorm.io/gorm"
)

func GetStudentsBy(studentId string,name string,classId string, pageSize int, pageIndex int) (res []*models.Student) {
	buildStudentQueryFrom(studentId,name,classId).
		Select("id", "student_id", "name", "class_id").
		Limit(pageSize).Offset(pageSize * (pageIndex - 1)).
		Find(&res)
	return
}

func GetStudentNumBy(studentId string,name string,classId string) (num int64){
	buildStudentQueryFrom(studentId,name,classId).Count(&num)
	return
}

func buildStudentQueryFrom(studentId string,name string,classId string) *gorm.DB {
	var filtered = db.Model(&models.Student{})
	if studentId != "" {
		filtered = filtered.Where("student_id LIKE ?", studentId+"%")
	}
	if name != "" {
		filtered = filtered.Where("name LIKE ?", "%"+name+"%")
	}
	if classId != ""{
		filtered = filtered.Where("class_id LIKE ?",classId+"%")
	}
	return filtered
}