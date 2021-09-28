package main

import "github.com/iyear/hduhelp-interview/server"

// @title 助手破冰API文档
// @version 1.0
// @description

// @contact.name iyear
// @contact.url https://github.com/iyear
// @contact.email ljyngup@gmail.com

// @host localhost:8080
// @BasePath /api/v1

// @tag.name photo
// @tag.description 图片前缀为 [host]/upload/

// @securityDefinitions.apikey authorization
// @in header
// @name authorization
func main() {
	server.Run()
}
