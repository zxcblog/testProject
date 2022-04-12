// Package api
// @Author:        asus
// @Description:   $
// @File:          address_controller
// @Data:          2022/4/1116:17
//
package api

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"new-project/pkg/app"
)

type AddressController struct {
	Ctx iris.Context
}

// Post 新增地址
// @Summary      新增地址
// @Description  新增地址
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Tags         用户地址
// @Success      200  {object}  app.Response{}
// @Router       /address [post]
func (this *AddressController) Post() *app.Response {

	fmt.Println("测试地址访问")

	return nil
}
