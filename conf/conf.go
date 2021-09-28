package conf

import (
	"github.com/iyear/hduhelp-interview/util"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var App *AppConf

func Init() {
	conf := viper.New()
	conf.SetConfigName("config")
	conf.AddConfigPath(".")

	if err := conf.ReadInConfig(); err != nil {
		zap.S().Fatalw("failed to read config",
			"error", err)
		return
	}

	if err := conf.Unmarshal(&App); err != nil {
		zap.S().Fatalw("failed to unmarshal app struct",
			"error", err)
		return
	}

	// 新建目录
	if err := util.CreatePathIfNotExists(App.Path.Photo); err != nil {
		zap.S().Fatalw("failed to make dir of photo",
			"error", err,
			"path", App.Path.Photo)
		return
	}
	if err := util.CreatePathIfNotExists(App.Path.Temp); err != nil {
		zap.S().Fatalw("failed to make dir of temp",
			"error", err,
			"path", App.Path.Temp)
		return
	}
}
