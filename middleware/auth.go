package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/guonaihong/gout"
	"github.com/iyear/hduhelp-interview/api"
	"github.com/iyear/hduhelp-interview/model"
	"github.com/tidwall/gjson"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		staffID := getValidate(c.GetHeader("authorization"))
		if staffID == -1 {
			c.Abort()
			c.Redirect(http.StatusFound, api.GetRedirectUrl())
			return
		}
		c.Set("staffID", staffID)
		c.Next()
	}
}

// getValidate 返回学号，错误为-1
func getValidate(authHeader string) int64 {
	var resp *model.AuthAccept
	err := gout.GET("https://api.hduhelp.com/oauth/token/validate").
		SetHeader(gout.H{
			"authorization": authHeader,
		}).BindJSON(&resp).Do()
	if err != nil {
		return -1
	}
	if resp.Error != 0 {
		return -1
	}
	id := gjson.Get(string(resp.Data), "staff_id").Int()
	return id
}
