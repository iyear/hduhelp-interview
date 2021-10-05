package srv_me

import (
	"errors"
	"github.com/iyear/hduhelp-interview/db"
	"github.com/iyear/hduhelp-interview/model"
	"github.com/iyear/hduhelp-interview/service/srv_depart"
	"github.com/iyear/hduhelp-interview/service/srv_photo"
	"github.com/iyear/hduhelp-interview/service/srv_stu"
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
	if req == nil {
		return errors.New("req is nil")
	}
	if stu, err = srv_stu.GetStudent(2, staffID); err != nil {
		return err
	}
	// photo是否存在
	if p, err = srv_photo.GetPhotoByFile(req.Photo); err != nil {
		return err
	}
	// depart是否存在
	if _, err = srv_depart.GetDepart(req.Depart); err != nil {
		return err
	}
	stu.Photo = p.ID
	stu.Depart = req.Depart
	stu.Show = req.Show

	if err = srv_stu.UpdateStudent(stu); err != nil {
		return err
	}
	return nil
}
