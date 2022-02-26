package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"new-chat/controller/api"
	config2 "new-chat/pkg/config"

	"github.com/iris-contrib/middleware/cors"
)

func Router() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}))

	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Party("/system").Handle(new(api.SystemController))
	})

	// iris.WithoutServerError(iris.ErrServerClosed) 忽略iris框架服务启动时的Listen的错误
	app.Run(iris.Addr(config2.GetService().Port), iris.WithoutServerError(iris.ErrServerClosed))
}
