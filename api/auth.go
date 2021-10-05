package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guonaihong/gout"
	"github.com/iyear/hduhelp-interview/conf"
	"github.com/iyear/hduhelp-interview/model"
	"github.com/iyear/hduhelp-interview/service/srv_stu"
	"github.com/iyear/hduhelp-interview/util"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"net/http"
)

func Auth(c *gin.Context) {
	var token string
	err := func() error {
		// 拿AccessToken
		var resp *model.AuthAccept
		err := gout.GET("https://api.hduhelp.com/oauth/token").
			SetQuery(gout.H{
				"client_id":     conf.App.Auth.ClientID,
				"client_secret": conf.App.Auth.ClientSecret,
				"grant_type":    "authorization_code",
				"code":          c.Query("code"),
				"state":         c.Query("state"),
			}).BindJSON(&resp).Do()
		if err != nil {
			return err
		}
		if resp.Error != 0 {
			return errors.New("wrong code")
		}

		// 拿信息写入
		token = gjson.Get(string(resp.Data), "access_token").String()
		err = gout.GET("https://api.hduhelp.com/base/person/info").
			SetHeader(gout.H{
				"authorization": "token " + token,
			}).BindJSON(&resp).Do()
		if err != nil {
			return err
		}
		if resp.Error != 0 {
			return errors.New("wrong code")
		}

		staffID := gjson.Get(string(resp.Data), "STAFFID").Int()
		staffName := gjson.Get(string(resp.Data), "STAFFNAME").String()

		if srv_stu.IsExistStudent(2, staffID) {
			return nil
		}
		err = srv_stu.AddStudent(&model.Student{
			StaffID:   staffID,
			StaffName: staffName,
			Photo:     -1, // 默认图片
			Depart:    -1, // 默认部门
			Show:      1,  // 默认显示
		})
		if err != nil {
			return err
		}
		return nil
	}()
	if err != nil {
		c.Redirect(http.StatusFound, GetRedirectUrl())
		zap.S().Errorw("failed to get auth",
			"error", err,
			"")
		return
	}
	c.Redirect(http.StatusFound, fmt.Sprintf("%s#/token=%s", conf.App.Url, token))
}
func GetRedirectUrl() string {
	s := util.GetRandomString(10)
	return fmt.Sprintf("https://api.hduhelp.com/oauth/authorize?response_type=code&client_id=%s&redirect_uri=%s&state=%s",
		conf.App.Auth.ClientID,
		conf.App.Auth.RedirectUrl,
		s)
}
