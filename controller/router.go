package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"new-chat/config"
	"new-chat/controller/api"

	"github.com/iris-contrib/middleware/cors"
)

func Router() {
	app := iris.Default()

	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}))

	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Party("/system").Handle(new(api.SystemController))
	})

	app.Run(iris.Addr(config.GetService().Port))
}
