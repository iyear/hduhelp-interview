package srv_stu

import (
	"errors"
	"github.com/iyear/hduhelp-interview/db"
	"github.com/iyear/hduhelp-interview/model"
	"gorm.io/gorm"
)

func AddStudent(s *model.Student) error {
	return db.Mysql.Create(&s).Error
}

// GetStudent tp=1为内部id,tp=2为学号
func GetStudent(tp int, oid int64) (*model.Student, error) {
	var s *model.Student
	ex := map[int]string{
		1: "id",
		2: "staff_id",
	}
	r := db.Mysql.Where(ex[tp]+" = ?", oid).Limit(1).First(&s)
	return s, r.Error
}
func IsExistStudent(tp int, oid int64) bool {
	_, err := GetStudent(tp, oid)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
func UpdateStudent(s *model.Student) error {
	return db.Mysql.Save(&s).Error
}
func GetAllStudents() ([]*model.Student, error) {
	var stu []*model.Student
	err := db.Mysql.Find(&stu).Error
	return stu, err

}
