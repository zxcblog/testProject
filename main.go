package main

import (
	"new-project/cache"
	"new-project/controller"
	"new-project/global"
	"new-project/models"
	"new-project/pkg/app"
	"new-project/pkg/config"
	"new-project/pkg/logger"
	"new-project/pkg/upload"
)

// @title           测试商城项目
// @version         1.0
// @description     自己学习的测试商城项目
// @termsOfService  http://1.14.127.213:19610

// @host      127.0.0.1:19610
// @BasePath  /
func main() {
	// 读取配置文件
	config.InitConfig("./config.yml")

	// 日志文件加载
	global.Logger = logger.NewLogger(config.GetLogger())
	global.AccessLog = logger.NewAccessLogger(config.GetLogger())

	// 数据库链接
	var err error
	if global.DB, err = models.NewDBEngine(config.GetDB(), models.Models...); err != nil {
		panic(err)
	}

	// redis链接
	if global.Redis, err = cache.InitRedis(config.GetRedis()); err != nil {
		panic(err)
	}

	// 加载参数验证器
	global.Validate = app.NewTranslationIns(app.WithLabelOption("label"), app.WithRulesOption(&app.Rules), app.WithRulesMsgOption(&app.RulesMsg))

	// 加载oss存储
	if global.Upload, err = upload.NewClient(config.GetOss()); err != nil {
		panic(err)
	}

	controller.Router()
}
