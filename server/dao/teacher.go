package dao

import (
    "github.com/gonearewe/EasyTesting/models"
    "github.com/gonearewe/EasyTesting/utils"
    "gorm.io/gorm"
)

func GetTeachersBy(teacherId string, name string, pageSize int, pageIndex int) (res []*models.Teacher, num int64) {
    err := db.Transaction(func(tx *gorm.DB) error {
        err := buildTeacherQueryFrom(tx, teacherId, name).
            Select("id", "teacher_id", "name", "is_admin").
            Limit(pageSize).Offset(pageSize * (pageIndex - 1)).
            Find(&res).Error
        if err != nil {
            return err
        }
        return buildTeacherQueryFrom(tx, teacherId, name).Count(&num).Error
    })
    utils.PanicWhen(err)
    return
}

func buildTeacherQueryFrom(tx *gorm.DB, teacherId string, name string) *gorm.DB {
    var filtered = tx.Model(&models.Teacher{})
    if teacherId != "" {
        filtered = filtered.Where("teacher_id LIKE ?", teacherId+"%")
    }
    if name != "" {
        filtered = filtered.Where("name LIKE ?", "%"+name+"%")
    }
    return filtered
}

func GetTeacherByTeacherId(teacherId string) *models.Teacher {
    var ret models.Teacher
    err := db.Find(&ret, "teacher_id = ?", teacherId).Error
    utils.PanicWhen(err)
    return &ret
}

func CreateTeachers(teachers []*models.Teacher) {
    utils.PanicWhen(db.Create(&teachers).Error)
}

func UpdateTeacherById(t *models.Teacher) {
    var filtered = db.Model(t).Where("id = ?", t.ID)
    var err error
    if t.Password == "" && t.Salt == "" {
        err = filtered.Updates(
            // Updates with map instead of struct to avoid fields of default value being ignored
            map[string]interface{}{"teacher_id": t.TeacherID, "name": t.Name, "is_admin": t.IsAdmin}).Error
    } else {
        err = filtered.Updates(
            map[string]interface{}{"teacher_id": t.TeacherID, "name": t.Name, "is_admin": t.IsAdmin,
                "password": t.Password, "salt": t.Salt}).Error
    }
    utils.PanicWhen(err)
}

func DeleteTeachers(ids []int) {
    deleteBy(ids,models.Teacher{})
}
