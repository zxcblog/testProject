package main

import (
	"new-chat/config"
	"new-chat/controller"
)

func main() {
	// 读取配置文件
	config.InitConfig("./config/config.yml")

	// 日志文件加载

	controller.Router()
}
