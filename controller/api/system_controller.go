package api

import (
	"github.com/kataras/iris/v12"
	"new-chat/pkg/app"
)

type SystemController struct {
	Ctx iris.Context
}

func (t *SystemController) GetPing() *app.Response {
	return app.ResponseMsg("pong")
}
