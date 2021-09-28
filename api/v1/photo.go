package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/iyear/hduhelp-interview/api"
	"github.com/iyear/hduhelp-interview/conf/e"
	"github.com/iyear/hduhelp-interview/model"
	"github.com/iyear/hduhelp-interview/service/srv_photo"
	"go.uber.org/zap"
	"mime/multipart"
	"net/http"
)

// GetStaffPhoto 获取当前照片
// @Summary 获取当前照片
// @Tags photo
// @Produce json
// @Param authorization header string true "token <AccessToken>"
// @Success 200 object api.Resp{data=model.Photo}
// @Router /photo/me [get]
func GetStaffPhoto(c *gin.Context) {
	staffID := c.GetInt64("staffID")
	p, err := srv_photo.GetStaffPhoto(staffID)
	if err != nil {
		api.RespFmt(c, http.StatusOK, &api.Resp{
			Error:    e.ERROR_GET_STAFF_PHOTO_FAIL,
			Msg:      e.GetMsg(e.ERROR_GET_STAFF_PHOTO_FAIL),
			Redirect: "",
			Data:     nil,
		})
		zap.S().Errorw("failed to get staff photo",
			"error", err,
			"staffID", staffID)
		return
	}
	api.RespFmt(c, http.StatusOK, &api.Resp{
		Error:    e.SUCCESS,
		Msg:      e.GetMsg(e.SUCCESS),
		Redirect: "",
		Data:     p,
	})
}

// UploadPhoto 上传照片
// @Summary 上传照片
// @Tags photo
// @Accept mpfd
// @Produce json
// @Param authorization header string true "token <AccessToken>"
// @Param photo formData file true "照片文件"
// @Success 200 object api.Resp{data=model.Photo}
// @Router /photo/upload [post]
func UploadPhoto(c *gin.Context) {

	p, err := func() (*model.Photo, error) {
		var (
			f   *multipart.FileHeader
			p   *model.Photo
			err error
		)
		if f, err = c.FormFile("photo"); err != nil {
			return nil, err
		}

		if p, err = srv_photo.UploadPhoto(f); err != nil {
			return nil, err
		}

		return p, nil
	}()
	if err != nil {
		api.RespFmt(c, http.StatusOK, &api.Resp{
			Error:    e.ERROR_UPLOAD_PHOTO_FAIL,
			Msg:      e.GetMsg(e.ERROR_UPLOAD_PHOTO_FAIL),
			Redirect: "",
			Data:     nil,
		})
		zap.S().Errorw("failed to save file",
			"error", err)
		return
	}
	api.RespFmt(c, http.StatusOK, &api.Resp{
		Error:    e.SUCCESS,
		Msg:      e.GetMsg(e.SUCCESS),
		Redirect: "",
		Data:     p,
	})

}
