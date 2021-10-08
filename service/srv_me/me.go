package srv_me

import (
	"github.com/iyear/hduhelp-interview/db"
	"github.com/iyear/hduhelp-interview/model"
)

func GetMe(staffID int64) (*model.GetMeResp, error) {
	var resp *model.GetMeResp
	if err := db.Mysql.Model(&model.Student{}).
		Select("students.staff_id, students.staff_name, students.show, photos.file AS photo, departs.name AS depart").
		Joins("LEFT JOIN photos ON photos.id = students.photo").
		Joins("LEFT JOIN departs ON departs.id = students.depart").
		Where("students.staff_id = ?", staffID).
		Limit(1).First(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}
func UpdateMe(staffID int64, req *model.UpdateMeReq) error {
	var (
		stu *model.Student
		p   *model.Photo
		err error
	)

	if err = db.Mysql.Where("staff_id = ?", staffID).Limit(1).First(&stu).Error; err != nil {
		return err
	}
	// photo是否存在
	if err = db.Mysql.Where("file = ?", req.Photo).Limit(1).First(&p).Error; err != nil {
		return err
	}
	// depart是否存在
	if err = db.Mysql.Where("id = ?", req.Depart).Limit(1).First(&model.Depart{}).Error; err != nil {
		return err
	}
	stu.Photo = p.ID
	stu.Depart = req.Depart
	stu.Show = req.Show

	if err = db.Mysql.Save(&stu).Error; err != nil {
		return err
	}
	return nil
}
