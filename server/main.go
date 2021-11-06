package main

import (
	"orderserve/config"
	"orderserve/router"
)

func main() {
	// 初始化配置
	config.ConfInit()
	// 启用数据库
	config.InitRrs()
	// 启用路由
	app := router.NewRouter()

	_ = app.Run("192.168.1.102:8090")
}

