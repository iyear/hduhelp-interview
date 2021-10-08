package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iyear/hduhelp-interview/conf"
	"github.com/iyear/hduhelp-interview/service/srv_auth"
	"github.com/iyear/hduhelp-interview/util"
	"go.uber.org/zap"
	"net/http"
)

func Auth(c *gin.Context) {
	token, err := srv_auth.Auth(c.Query("code"), c.Query("state"))
	if err != nil {
		c.Redirect(http.StatusFound, GetRedirectUrl())
		zap.S().Errorw("failed to get auth",
			"error", err)
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
