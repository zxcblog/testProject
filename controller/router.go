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

	// 部分请求使用options协议先进行访问，并在请求头中说明对应的请求方法
	// 当对应的options请求在路由中没有监听时， 会将路由请求返回并跳转路由
	// 为了防止这种情况， 使用默认路径来进行缓冲，让iris框架获取请求中的实际请求协议
	app.Any("{root:path}", func(ctx iris.Context) {
		// TODO 记录访问不存在的路径以及请求的用户信息
		ctx.StatusCode(404)
	})

	// iris.WithoutServerError(iris.ErrServerClosed) 忽略iris框架服务启动时的Listen的错误
	// iris.WithOptimizations 应用程序会尽可能优化以获得最佳性能
	app.Run(iris.Addr(":19610"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
