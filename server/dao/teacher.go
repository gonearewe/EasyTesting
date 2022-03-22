package dao

import (
    "github.com/gonearewe/EasyTesting/models"
    "github.com/gonearewe/EasyTesting/utils"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

func GetTeachersBy(teacherId string, name string, pageSize int, pageIndex int) (ret []*models.Teacher, num int64) {
    err := db.Transaction(func(tx *gorm.DB) error {
        err := buildTeacherQueryFrom(tx, teacherId, name).
            // `password` and `salt` field is excluded because clients don't need it.
            Select("id", "teacher_id", "name", "is_admin").
            Limit(pageSize).Offset(pageSize * (pageIndex - 1)).
            Find(&ret).Error
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
    return filtered.Order("id")
}

// GetTeacherByTeacherId gets a teacher by the teacher id including `password` and `salt` field.
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
    var filtered = db.Model(t)
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

func UpdateTeacherProfileById(t *models.Teacher) {
    var filtered = db.Model(t)
    var err error
    if t.Password == "" && t.Salt == "" {
        err = filtered.Updates(
            // Updates with map instead of struct to avoid fields of default value being ignored
            map[string]interface{}{"teacher_id": t.TeacherID, "name": t.Name}).Error
    } else {
        err = filtered.Updates(
            map[string]interface{}{"teacher_id": t.TeacherID, "name": t.Name,
                "password": t.Password, "salt": t.Salt}).Error
    }
    utils.PanicWhen(err)
}

func DeleteTeachers(ids []int) {
    err := db.Transaction(func(tx *gorm.DB) error {
        for _, id := range ids {
            // SELECT FOR UPDATE, make sure all the ids exist
            err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
                Select("id").Where("id = ?", id).First(&models.Teacher{}).Error
            if err != nil {
                return err
            }
        }
        // batch delete
        return tx.Delete(&models.Teacher{}, ids).Error
    })
    utils.PanicWhen(err)
}
