package srv_auth

import (
	"github.com/guonaihong/gout"
	"github.com/iyear/hduhelp-interview/conf"
	"github.com/iyear/hduhelp-interview/db"
	"github.com/iyear/hduhelp-interview/model"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

func Auth(code, state string) (string, error) {
	var (
		resp  *model.AuthAccept
		token string
	)
	err := gout.GET("https://api.hduhelp.com/oauth/token").
		SetQuery(gout.H{
			"client_id":     conf.App.Auth.ClientID,
			"client_secret": conf.App.Auth.ClientSecret,
			"grant_type":    "authorization_code",
			"code":          code,
			"state":         state,
		}).BindJSON(&resp).Do()
	if err != nil {
		return "", err
	}
	if resp.Error != 0 {
		return "", errors.New("wrong code")
	}

	// 拿信息写入
	token = gjson.Get(string(resp.Data), "access_token").String()
	err = gout.GET("https://api.hduhelp.com/base/person/info").
		SetHeader(gout.H{
			"authorization": "token " + token,
		}).BindJSON(&resp).Do()
	if err != nil {
		return "", err
	}
	if resp.Error != 0 {
		return "", errors.New("wrong code")
	}

	staffID := gjson.Get(string(resp.Data), "STAFFID").Int()
	staffName := gjson.Get(string(resp.Data), "STAFFNAME").String()

	if count := db.Mysql.Where("staff_id = ?", staffID).First(&model.Student{}).RowsAffected; count > 0 {
		return token, nil
	}

	if err = db.Mysql.Create(&model.Student{
		StaffID:   staffID,
		StaffName: staffName,
		Photo:     -1, // 默认图片
		Depart:    -1, // 默认部门
		Show:      1,  // 默认显示
	}).Error; err != nil {
		return "", err
	}
	return token, nil
}
