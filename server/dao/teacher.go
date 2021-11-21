package dao

import (
    "github.com/gin-gonic/gin"
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
    db.Find(&ret,"teacher_id = ?",teacherId)
    return &ret
}

func GetTeachersByTeacherId(pageSize int, pageIndex int) (res []models.Teacher) {
    db.Limit(pageSize).Offset(pageSize * pageIndex).Order("teacher_id").Find(&res)
    return
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

func DeleteTeacherByTeacherId(teacherId string)  {
    db.Delete(&models.Teacher{},"teacher_id = ?",teacherId)
}
