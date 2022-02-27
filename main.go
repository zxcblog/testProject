package main

import (
	"new-project/controller"
	"new-project/global"
	"new-project/models"
	"new-project/pkg/app"
	config2 "new-project/pkg/config"
	"new-project/pkg/logger"
)

// @title 测试商城项目
// @version 1.0
// @description 自己学习的测试商城项目
// @termsOfService http://127.0.0.1:19610

// @host 127.0.0.1:19610
// @BasePath /
func main() {
	// 读取配置文件
	config2.InitConfig("./config.yml")

	// 日志文件加载
	global.Logger = logger.NewLogger(config2.GetLogger())
	global.AccessLog = logger.NewAccessLogger(config2.GetLogger())

	// 数据库链接
	var err error
	if global.DB, err = models.NewDBEngine(config2.GetDB(), models.Models...); err != nil {
		panic(err)
	}

	// 加载参数验证器
	global.Validate = app.NewTranslationIns(app.WithLabelOption("label"), app.WithRulesOption(&app.Rules), app.WithRulesMsgOption(&app.RulesMsg))

	controller.Router()
}
