package controller

import (
	"fmt"
	"net/http"
	"new-project/controller/admin"
	"new-project/controller/api"
	"new-project/controller/comm"
	"new-project/controller/middleware"
	_ "new-project/docs"
	"new-project/pkg/config"

	"github.com/iris-contrib/swagger/v12"
	"github.com/kataras/iris/v12"

	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12/mvc"

	"github.com/iris-contrib/middleware/cors"
)

func Router() {
	app := iris.New()
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
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

	mvc.Configure(app, func(m *mvc.Application) {
		m.Router.Use(middleware.AccessLog) // 访问日志中间件
		//m.Router.Use(middleware.Limiter)   // 限流
		// 前台管理
		apiRoute := m.Party("/api")
		apiRoute.Party("/system").Handle(new(api.SystemController))
		apiRoute.Party("/user").Handle(new(api.UserController))       //用户
		apiRoute.Party("/captcha").Handle(new(api.CaptchaController)) //验证码

		// 后台管理
		adminRoute := m.Party("/admin")
		adminRoute.Party("/category").Handle(new(admin.CategoryController)) //分类
		adminRoute.Party("/brand").Handle(new(admin.BrandController))       //品牌
		adminRoute.Party("/product").Handle(new(admin.ProductController))   //商品

		// 通用路由
		commRoute := m.Party("/comm")
		commRoute.Party("/upload").Handle(new(comm.UploadController))

	})

	// iris.WithoutServerError(iris.ErrServerClosed) 忽略iris框架服务启动时的Listen的错误
	// iris.WithOptimizations 应用程序会尽可能优化以获得最佳性能
	app.Run(iris.Addr(":19610"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
