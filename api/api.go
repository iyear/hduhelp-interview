package api

import (
	"github.com/gin-gonic/gin"
	"github.com/iyear/hduhelp-interview/conf/e"
	"net/http"
)

type Resp struct {
	Error    int         `json:"error" example:"0"`
	Msg      string      `json:"msg" example:"ok"`
	Redirect string      `json:"redirect" example:" "`
	Data     interface{} `json:"data,omitempty"`
}

func RespFmt(c *gin.Context, code int, resp *Resp) {
	c.JSON(code, resp)
}
func ParamsCheck(c *gin.Context, v interface{}) error {
	err := c.ShouldBind(v)
	if err != nil {
		RespFmt(c, http.StatusOK, &Resp{
			Error:    e.INVALID_PARAMS,
			Msg:      err.Error(),
			Redirect: "",
			Data:     nil,
		})
		return err
	}
	return nil
}
