// Package db 省略dao层，直接在 service 里写数据操作
package db

import (
	"gorm.io/gorm"
)

var (
	Mysql *gorm.DB
)

func Init() {
	Mysql = initMysql()
}
