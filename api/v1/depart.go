package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iyear/hduhelp-interview/api"
	"github.com/iyear/hduhelp-interview/conf/e"
	"github.com/iyear/hduhelp-interview/model"
	"github.com/iyear/hduhelp-interview/service/srv_depart"
	"go.uber.org/zap"
	"net/http"
)

// GetAllDeparts 获取所有部门信息
// @Summary 获取所有部门信息
// @Tags depart
// @Produce json
// @Param authorization header string true "token <AccessToken>"
// @Success 200 object api.Resp{data=model.GetAllDepartsResp}
// @Router /depart/getAll [get]
func GetAllDeparts(c *gin.Context) {
	var (
		departs []*model.Depart
		err     error
	)
	if departs, err = srv_depart.GetAllDeparts(); err != nil {
		api.RespFmt(c, http.StatusOK, &api.Resp{
			Error:    e.ERROR_GET_ALL_DEPARTS_FAIL,
			Msg:      e.GetMsg(e.ERROR_GET_ALL_DEPARTS_FAIL),
			Redirect: "",
			Data:     nil,
		})
		zap.S().Errorw("failed to get all departs",
			"error", err)
		return
	}

	api.RespFmt(c, http.StatusOK, &api.Resp{
		Error:    0,
		Msg:      e.GetMsg(e.SUCCESS),
		Redirect: "",
		Data: &model.GetAllDepartsResp{
			Count:   len(departs),
			Departs: departs,
		},
	})
}
