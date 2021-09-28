package db

import (
	"fmt"
	"github.com/iyear/hduhelp-interview/conf"
	"github.com/iyear/hduhelp-interview/model"
	"github.com/iyear/hduhelp-interview/util"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func initMysql() *gorm.DB {
	dial := mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.App.Mysql.User,
		conf.App.Mysql.Password,
		conf.App.Mysql.Host,
		conf.App.Mysql.Port,
		conf.App.Mysql.Database,
	))

	d, err := gorm.Open(dial, &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		Logger: logger.Default.LogMode(util.IF(conf.App.Debug, logger.Info, logger.Silent).(logger.LogLevel)),
	})
	if err != nil {
		zap.S().Fatalw("failed to open sqlite db",
			"error", err,
			"config", conf.App.Mysql)
		return nil
	}

	// 迁移
	if err = d.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(&model.Student{}, &model.Depart{}, &model.Photo{}); err != nil {
		zap.S().Fatalw("failed to migrate data",
			"error", err)
		return nil
	}

	return d
}
