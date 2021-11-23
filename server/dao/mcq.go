package dao

import "github.com/gonearewe/EasyTesting/models"

func GetMcqsBy(teacherId string, pageSize int, pageIndex int) (ret []*models.Mcq) {
    if teacherId != "" {
        db.Limit(pageSize).Offset(pageSize*pageIndex).Find(&ret, "teacher_id LIKE ?", teacherId+"%")
    } else {
        db.Limit(pageSize).Offset(pageSize*pageIndex).Find(&ret)
    }
    return
}

func CreateMcqs(questions []*models.Mcq) {
    db.Create(&questions)
}

func UpdateMcqs(questions []*models.Mcq) {
    db.Updates(&questions)
}

func DeleteMcqs(ids []int)  {
    db.Delete(ids)
}
