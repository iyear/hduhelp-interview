package srv_option

import (
	"errors"
	"github.com/iyear/hduhelp-interview/db"
	"github.com/iyear/hduhelp-interview/model"
	"github.com/iyear/hduhelp-interview/service/srv_photo"
	"github.com/iyear/hduhelp-interview/service/srv_stu"
	"math/rand"
	"time"
)

func GetOptions(n int, depart int64) (*model.GetOptionsResp, error) {
	var (
		results []*model.Option
		stu     []*model.Student
		shows   []*model.Student
		photo   *model.Photo
		err     error
	)
	rand.Seed(time.Now().UnixNano())

	if stu, err = srv_stu.GetAllStudents(); err != nil {
		return nil, err
	}
	if len(stu) < n {
		return nil, errors.New("n is bigger than len of students")
	}

	if shows = getFilterStu(stu, depart); len(shows) < n {
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

	if photo, err = srv_photo.GetPhotoByID(shows[rand.Intn(n)].Photo); err != nil {
		return nil, err
	}
	return &model.GetOptionsResp{
		Photo:   photo.File,
		Options: results,
	}, nil
}
func getFilterStu(stu []*model.Student, depart int64) []*model.Student {
	var shows []*model.Student
	for _, s := range stu {
		// depart = -1 不筛选部门
		if s.Show == 1 && (s.Depart == depart || depart == -1) {
			shows = append(shows, s)
		}
	}
	return shows
}

// JudgeOption photo为图片file,id为 GetOptions 的返回值中的id
func JudgeOption(photo string, id int64) (*model.JudgeOptionResp, error) {
	var (
		choose  *model.Student
		right   []*model.Student
		results []string
		p       *model.Photo
		err     error
	)
	if p, err = srv_photo.GetPhotoByFile(photo); err != nil {
		return nil, err
	}
	if choose, err = srv_stu.GetStudent(1, id); err != nil {
		return nil, err
	}

	if choose.Photo == p.ID {
		return &model.JudgeOptionResp{
			Right: 1,
		}, nil
	}

	// 找到正确的照片，有可能多人用同一张默认照片，返回数组
	if err = db.Mysql.Where("photo = ?", p.ID).Find(&right).Error; err != nil {
		return nil, err
	}
	for _, s := range right {
		results = append(results, s.StaffName)
	}
	return &model.JudgeOptionResp{
		Right: 2,
		Name:  results,
	}, nil
}
