// Package admin
// @Author:        asus
// @Description:   $
// @File:          category_controller
// @Data:          2022/2/2710:44
//
package admin

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"new-project/global"
	"new-project/pkg/app"
)

type CategoryController struct {
	Ctx iris.Context
}

// 创建分类
type CategoryControllerPost struct {
	CategoryName string `validate:"required,min=1,max=20" label:"分类名称" json:"categoryName"` // 分类名称
	CategoryID   uint   `validate:"-" label:"所属分类" json:"categoryID" default:"0"`           // 所属分类
	IsFinal      bool   `validate:"-" label:"是否为最终类" json:"isFinal"  default:"false"`       // 是否为最终类
}

// @Summary 添加商品分类
// @Description 后台管理人员添加商品分类
// @Accept json
// @Produce json
// @param root body CategoryControllerPost true "添加商品分类"
// @Tags 商品分类
// @Success 200 {object} app.Response
// @Router /admin/category [post]
func (c *CategoryController) Post() *app.Response {
	param := &CategoryControllerPost{}
	err := c.Ctx.ReadJSON(param)
	fmt.Println("打印接收值", param, err)

	err = global.Validate.ValidateParam(param)
	fmt.Println(err)

	return app.ResponseMsg("创建成功")
}
