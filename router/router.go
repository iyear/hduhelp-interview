package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iyear/hduhelp-interview/api"
	v1 "github.com/iyear/hduhelp-interview/api/v1"
	"github.com/iyear/hduhelp-interview/conf"
	_ "github.com/iyear/hduhelp-interview/docs"
	"github.com/iyear/hduhelp-interview/middleware"
	"github.com/iyear/hduhelp-interview/util"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Init() *gin.Engine {
	gin.SetMode(util.IF(conf.App.Debug, gin.DebugMode, gin.ReleaseMode).(string))

	r := gin.New()
	r.Use(middleware.Recovery(), gin.Logger())

	r.StaticFS("/upload", gin.Dir(conf.App.Path.Photo, false))

	r.GET("/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth", api.Auth)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.RateLimit(middleware.RateLimitInit()), middleware.Auth())
	{
		apiV1.GET("/option/get", v1.GetOptions)
		apiV1.GET("/option/judge", v1.JudgeOption)

		apiV1.GET("/me/info", v1.GetMe)
		apiV1.POST("/me/update", v1.UpdateMe)

		apiV1.POST("/photo/upload", v1.UploadPhoto)
		apiV1.GET("/photo/me", v1.GetStaffPhoto)

		apiV1.GET("/depart/getAll", v1.GetAllDeparts)
	}
	return r
}
