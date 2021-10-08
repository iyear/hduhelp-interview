package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iyear/hduhelp-interview/api"
	"github.com/iyear/hduhelp-interview/conf/e"
	"github.com/iyear/hduhelp-interview/model"
	"github.com/iyear/hduhelp-interview/service/srv_option"
	"go.uber.org/zap"
	"net/http"
)

// GetOptions 获取选项
// @Summary 获取选项
// @Tags option
// @Produce json
// @Param authorization header string true "token <AccessToken>"
// @Param num query int true "数量，大于总数或显示照片人数将返回错误"
// @Param depart query int true "筛选的部门,-2为不限"
// @Success 200 object api.Resp{data=model.GetOptionsResp}
// @Router /option/get [get]
func GetOptions(c *gin.Context) {
	var (
		req  *model.GetOptionsReq
		resp *model.GetOptionsResp
		err  error
	)

	err = api.ParamsCheck(c, &req)
	if err != nil {
		return
	}

	if resp, err = srv_option.GetOptions(req.Num, req.Depart); err != nil {
		api.RespFmt(c, http.StatusOK, &api.Resp{
			Error:    e.ERROR_GET_OPTIONS_FAIL,
			Msg:      e.GetMsg(e.ERROR_GET_OPTIONS_FAIL),
			Redirect: "",
			Data:     nil,
		})
		zap.S().Errorw("failed to get options",
			"error", err,
			"params", req)
		return
	}
	api.RespFmt(c, http.StatusOK, &api.Resp{
		Error:    e.SUCCESS,
		Msg:      e.GetMsg(e.SUCCESS),
		Redirect: "",
		Data:     resp,
	})
}

// JudgeOption 判断选项
// @Summary 判断选项
// @Tags option
// @Produce json
// @Param authorization header string true "token <AccessToken>"
// @Param photo query string true "照片file"
// @Param id query int true "学生id"
// @Success 200 object api.Resp{data=model.JudgeOptionResp}
// @Router /option/judge [get]
func JudgeOption(c *gin.Context) {
	var (
		req  *model.JudgeOptionReq
		resp *model.JudgeOptionResp
		err  error
	)
	if err = api.ParamsCheck(c, &req); err != nil {
		return
	}
	if resp, err = srv_option.JudgeOption(req.Photo, req.ID); err != nil {
		api.RespFmt(c, http.StatusOK, &api.Resp{
			Error:    e.ERROR_JUDGE_OPTION_FAIL,
			Msg:      e.GetMsg(e.ERROR_JUDGE_OPTION_FAIL),
			Redirect: "",
			Data:     nil,
		})
		zap.S().Errorw("failed to judge option",
			"error", err,
			"params", req)
		return
	}
	api.RespFmt(c, http.StatusOK, &api.Resp{
		Error:    e.SUCCESS,
		Msg:      e.GetMsg(e.SUCCESS),
		Redirect: "",
		Data:     resp,
	})

}
