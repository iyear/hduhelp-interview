package srv_option

import (
	"errors"
	"github.com/iyear/hduhelp-interview/db"
	"github.com/iyear/hduhelp-interview/model"
	"math/rand"
	"time"
)

func GetOptions(n int, depart int64) (*model.GetOptionsResp, error) {
	var (
		results []*model.Option
		shows   []*model.Student
		photo   *model.Photo
		err     error
	)
	rand.Seed(time.Now().UnixNano())
	// 不限部门
	if depart == -2 {
		err = db.Mysql.Where("`show` = ?", 1).Find(&shows).Error
	} else {
		err = db.Mysql.Where("`show` = ? and depart = ?", 1, depart).Find(&shows).Error
	}

	if err != nil {
		return nil, err
	}
	if len(shows) < n {
		return nil, errors.New("n is bigger than len of shows")
	}

	rand.Shuffle(len(shows), func(i, j int) {
		shows[i], shows[j] = shows[j], shows[i]
	})
	for i := 0; i < n; i++ {
		results = append(results, &model.Option{
			ID:   shows[i].ID,
			Name: shows[i].StaffName,
		})
	}

	if err = db.Mysql.Where("id = ?", shows[rand.Intn(n)].Photo).
		Limit(1).
		First(&photo).Error; err != nil {
		return nil, err
	}
	return &model.GetOptionsResp{
		Photo:   photo.File,
		Options: results,
	}, nil
}

// JudgeOption photo为图片file,id为 GetOptions 的返回值中的id
func JudgeOption(photo string, id int64) (*model.JudgeOptionResp, error) {
	err := db.Mysql.
		Joins("LEFT JOIN photos ON students.photo = photos.id").
		Where("students.id = ? AND photos.file = ?", id, photo).
		Limit(1).First(&model.Student{}).Error
	if err == nil {
		return &model.JudgeOptionResp{
			Right: 1,
		}, nil
	}

	// 找到正确的照片，有可能多人用同一张默认照片，返回数组
	var names []*struct {
		Name string `json:"name"`
	}
	err = db.Mysql.
		Model(&model.Student{}).
		Select("students.staff_name AS name").
		Joins("LEFT JOIN photos ON photos.id = students.photo").
		Where("photos.file = ?", photo).
		Find(&names).Error
	if err != nil {
		return nil, err
	}

	var results []string
	for _, name := range names {
		results = append(results, name.Name)
	}
	return &model.JudgeOptionResp{
		Right: 2,
		Name:  results,
	}, nil
}
