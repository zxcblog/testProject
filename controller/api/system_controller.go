package api

import (
	"github.com/kataras/iris/v12"
	"new-project/pkg/app"
)

type SystemController struct {
	Ctx iris.Context
}

// @Summary 测试系统接口
// @Description 测试系统接口是否能正常访问
// @Accept json
// @Produce json
// @Tags system
// @Success 200 {object} app.Response pong
// @Router /api/system/ping [get]
func (t *SystemController) GetPing() *app.Response {
	return app.ResponseMsg("pong")
}
