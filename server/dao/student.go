package dao

import (
	"github.com/gonearewe/EasyTesting/models"
	"github.com/gonearewe/EasyTesting/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GetStudentsBy searches the database for Student whose studentId or classId (string) starts
// with respective param and name contains `name`, it only returns records in given `pageIndex` (
// 1-based) in the increasing order of id, but plus the total number of all filtered.
// When any error occurs, it panics.
func GetStudentsBy(studentId string, name string, classId string,
	pageSize int, pageIndex int) (res []*models.Student, num int64) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := buildStudentQueryFrom(tx, studentId, name, classId).
			Select("id", "student_id", "name", "class_id").
			Limit(pageSize).Offset(pageSize * (pageIndex - 1)).
			Find(&res).Error
		if err != nil {
			return err
		}
		return buildStudentQueryFrom(tx, studentId, name, classId).Count(&num).Error
	})
	utils.PanicWhen(err)
	return
}

func GetStudentBy(studentId string,name string) *models.Student {
	var ret models.Student
	err := db.Select("student_id", "name", "class_id").
		Where("student_id = ? AND name = ?", studentId,name).First(&ret).Error
	utils.PanicWhen(err)
	return &ret
}

func buildStudentQueryFrom(tx *gorm.DB, studentId string, name string, classId string) *gorm.DB {
	var filtered = tx.Model(&models.Student{})
	if studentId != "" {
		filtered = filtered.Where("student_id LIKE ?", studentId+"%")
	}
	if name != "" {
		filtered = filtered.Where("name LIKE ?", "%"+name+"%")
	}
	if classId != "" {
		filtered = filtered.Where("class_id LIKE ?", classId+"%")
	}
	return filtered
}

// CreateStudents stores given Student into the database,
// with their id property ignored and handled by the database. When any error occurs,
// it panics and none of the given student will be created alone.
func CreateStudents(students []*models.Student) {
	utils.PanicWhen(db.Create(&students).Error)
}

// UpdateStudentById updates all the columns of given Student,
// the record to be updated will be specified by given student's id.
// When any error occurs, it panics and the given student will not be updated.
func UpdateStudentById(t *models.Student) {
	err := db.Transaction(func(tx *gorm.DB) error {
		var err error
		err = db.Where("id = ?", t.ID).Updates(t).Error
		if err != nil {
			return err
		}

		return db.Model(&models.ExamSession{}).Where("student_id = ?", t.StudentID).
			Update("student_name", t.Name).Error
	})
	utils.PanicWhen(err)
}

// DeleteStudents deletes all the records whose id is in given `ids`.
// If any id in given `ids` doesn't exist, it refuses to proceed and throws an error.
// When any error occurs, it panics and none of the given student will be deleted alone.
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
