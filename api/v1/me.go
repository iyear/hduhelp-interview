package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iyear/hduhelp-interview/api"
	"github.com/iyear/hduhelp-interview/conf/e"
	"github.com/iyear/hduhelp-interview/model"
	"github.com/iyear/hduhelp-interview/service/srv_me"
	"go.uber.org/zap"
	"net/http"
)

// GetMe 获取个人信息
// @Summary 获取个人信息
// @Tags me
// @Produce json
// @Param authorization header string true "token <AccessToken>"
// @Success 200 object api.Resp{data=model.GetMeResp}
// @Router /me/info [get]
func GetMe(c *gin.Context) {
	s, err := srv_me.GetMe(c.GetInt64("staffID"))
	if err != nil {
		api.RespFmt(c, http.StatusOK, &api.Resp{
			Error:    e.ERROR_GET_ME_FAIL,
			Msg:      e.GetMsg(e.ERROR_GET_ME_FAIL),
			Redirect: "",
			Data:     nil,
		})
		zap.S().Errorw("failed to get student",
			"error", err,
			"stu", s)
		return
	}

	api.RespFmt(c, http.StatusOK, &api.Resp{
		Error:    e.SUCCESS,
		Msg:      e.GetMsg(e.SUCCESS),
		Redirect: "",
		Data:     s,
	})
}

// UpdateMe 更新个人信息
// @Summary 更新个人信息
// @Tags me
// @Accept json
// @Produce json
// @Param authorization header string true "token <AccessToken>"
// @Param object body model.UpdateMeReq true "更改后的信息"
// @Success 200 object api.Resp
// @Router /me/update [post]
func UpdateMe(c *gin.Context) {
	var req *model.UpdateMeReq

	if err := api.ParamsCheck(c, &req); err != nil {
		return
	}

	staffID := c.GetInt64("staffID")
	if err := srv_me.UpdateMe(staffID, req); err != nil {
		api.RespFmt(c, http.StatusOK, &api.Resp{
			Error:    e.ERROR_UPDATE_ME_FAIL,
			Msg:      e.GetMsg(e.ERROR_UPDATE_ME_FAIL),
			Redirect: "",
			Data:     nil,
		})
		zap.S().Errorw("failed to update me",
			"error", err,
			"params", req,
			"staffID", staffID)
		return
	}
	api.RespFmt(c, http.StatusOK, &api.Resp{
		Error:    e.SUCCESS,
		Msg:      e.GetMsg(e.SUCCESS),
		Redirect: "",
		Data:     nil,
	})
}
