package dao

import (
	"github.com/gonearewe/EasyTesting/models"
	"github.com/gonearewe/EasyTesting/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func CreateStudents(students []*models.Student) {
	db.Create(&students)
}

func UpdateStudentById(t *models.Student) {
	err := db.Where("id = ?", t.ID).Updates(t).Error
	utils.PanicWhen(err)
}

func DeleteStudents(ids []int) {
	err := db.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			// SELECT FOR UPDATE, make sure all the ids exist
			err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				Select("id").Where("id = ?", id).First(&models.Student{}).Error
			if err != nil {
				return err
			}
		}
		// batch delete
		return tx.Delete(&models.Student{}, ids).Error
	})
	utils.PanicWhen(err)
}