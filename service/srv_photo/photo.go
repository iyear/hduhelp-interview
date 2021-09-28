package srv_photo

import (
	"errors"
	"github.com/google/uuid"
	"github.com/iyear/hduhelp-interview/conf"
	"github.com/iyear/hduhelp-interview/db"
	"github.com/iyear/hduhelp-interview/model"
	"github.com/iyear/hduhelp-interview/service/srv_stu"
	"github.com/iyear/hduhelp-interview/util"
	"mime/multipart"
	"strings"
)

func GetStaffPhoto(staffID int64) (*model.Photo, error) {
	var (
		p   *model.Photo
		stu *model.Student
		err error
	)
	stu, err = srv_stu.GetStudent(2, staffID)
	if err != nil {
		return nil, err
	}
	p, err = GetPhotoByID(stu.Photo)
	if err != nil {
		return nil, err
	}
	return p, nil
}
func GetPhotoByID(id int64) (*model.Photo, error) {
	var p *model.Photo
	if err := db.Mysql.Where("id = ?", id).Limit(1).First(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}
func GetPhotoByFile(file string) (*model.Photo, error) {
	var p *model.Photo
	if err := db.Mysql.Where("file = ?", file).Limit(1).First(&p).Error; err != nil {
		return nil, err
	}
	return p, nil
}
func UploadPhoto(file *multipart.FileHeader) (*model.Photo, error) {
	var name string
	// 大小验证
	if conf.App.Photo.Min*1024 > file.Size || conf.App.Photo.Max*1024 < file.Size {
		return nil, errors.New("wrong photo size")
	}
	// 校验mime
	contentType, err := util.GetMPFDContentType(file)
	if err != nil {
		return nil, err
	}
	if !util.IsInStringSlice(contentType, conf.App.Photo.MIME) {
		return nil, errors.New("wrong photo format")
	}

	// 随机文件名
	sep := strings.Split(contentType, "/")
	ext := "." + sep[1]

	for {
		name = uuid.New().String() + ext
		if !util.IsExists(conf.App.Path.Photo + name) {
			break
		}
	}
	if err = util.SaveUploadedFile(file, conf.App.Path.Photo+name); err != nil {
		return nil, err
	}
	var p = &model.Photo{
		File: name,
		Size: file.Size,
	}
	r := db.Mysql.Create(&p)
	return p, r.Error
}
