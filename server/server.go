package server

import (
	"fmt"
	"github.com/iyear/hduhelp-interview/conf"
	"github.com/iyear/hduhelp-interview/db"
	"github.com/iyear/hduhelp-interview/logger"
	"github.com/iyear/hduhelp-interview/router"
	"go.uber.org/zap"
)

func Run() {
	logger.Init()
	conf.Init()
	db.Init()
	e := router.Init()

	if err := e.Run(fmt.Sprintf(":%d", conf.App.Port)); err != nil {
		zap.S().Fatalw("failed to run gin engine",
			"error", err)
		return
	}
}
