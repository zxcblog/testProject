package controller

import (
	"fmt"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"new-project/controller/admin"
	"new-project/controller/api"
	_ "new-project/docs"
	"new-project/pkg/config"

	"github.com/iris-contrib/middleware/cors"
)

func Router() {
	app := iris.New()
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            config.GetService().DebugMode,
	}))

	// swagger文档地址，只有在debug模式下才能访问
	if config.GetService().DebugMode {
		app.Logger().SetLevel("debug") // 设置调式模式
		swaggerUI := swagger.WrapHandler(swaggerFiles.Handler,
			swagger.URL(fmt.Sprintf("%s/swagger/doc.json", config.GetService().Host)),
			swagger.DeepLinking(true),
		)
		app.Get("/swagger/{any:path}", swaggerUI)
		app.Get("/swagger", swaggerUI)
	}

	// 系统路由
	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Party("/system").Handle(new(api.SystemController))
	})

	// 管理员操作
	mvc.Configure(app.Party("/admin"), func(m *mvc.Application) {
		m.Party("/category").Handle(new(admin.CategoryController))
	})

	// iris.WithoutServerError(iris.ErrServerClosed) 忽略iris框架服务启动时的Listen的错误
	app.Run(iris.Addr(config.GetService().Port), iris.WithoutServerError(iris.ErrServerClosed))
}
