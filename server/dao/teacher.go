package dao

import (
    "github.com/gin-gonic/gin"
    "github.com/gonearewe/EasyTesting/models"
    "github.com/gonearewe/EasyTesting/utils"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

func GetTeachersBy(teacherId string, name string, pageSize int, pageIndex int) (res []*models.Teacher) {
    buildTeacherQueryFrom(teacherId , name ).
        Select("id", "teacher_id", "name", "is_admin").
        Limit(pageSize).Offset(pageSize * (pageIndex - 1)).
        Find(&res)
    return
}

func GetTeacherNumBy(teacherId string, name string) (num int64){
    buildTeacherQueryFrom(teacherId , name ).Count(&num)
    return
}

func  buildTeacherQueryFrom(teacherId string, name string) *gorm.DB {
    var filtered = db.Model(&models.Teacher{})
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
    db.Find(&ret,"teacher_id = ?",teacherId)
    return &ret
}

func CreateTeachers(teachers []*models.Teacher) {
    db.Create(&teachers)
}

func UpdateTeacherByTeacherId(t *models.Teacher) {
    var filtered = db.Where("teacher_id = ?", t.TeacherID)
    if t.Password == "" && t.Salt == ""{
        filtered.Updates(gin.H{"name": t.Name, "is_admin": t.IsAdmin})
    }else {
        filtered.Updates(gin.H{"name": t.Name, "password": t.Password, "salt": t.Salt, "is_admin": t.IsAdmin})
    }
}

func DeleteTeachers(ids []int)  {
    err := db.Transaction(func(tx *gorm.DB) error {
        tmpTeacher := &models.Teacher{}
        for _, id := range ids {
            // SELECT FOR UPDATE, make sure all the ids exist
            err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
                Select("id").Where("id = ?", id).First(tmpTeacher).Error
            if err != nil {
                return err
            }
        }
        // batch delete
        return tx.Delete(&models.Teacher{}, ids).Error
    })
    utils.PanicWhen(err)
}
