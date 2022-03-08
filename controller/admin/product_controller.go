package admin

import (
	"new-project/global"
	"new-project/pkg/app"

	"github.com/kataras/iris/v12"
)

type ProductController struct {
	Ctx iris.Context
}

type Attr struct {
	AttrName      string `validate:"required"`
	AttrGroupName string `validate:"required"`
}

//添加商品校验
type PostProductAddCheckedParams struct {
	ProductName string `validata:"required,max=100" label:"商品名称" json:"productName"` //商品名称
}

//商品添加
func (p *ProductController) PostProductadd() *app.Response {
	params := PostProductAddCheckedParams{}

	if err := p.Ctx.ReadJSON(params); err != nil {
		return app.ResponseMsg(err.Error())
	}

	if err := global.Validate.ValidateParam(params); err != nil {
		return app.ResponseMsg(err.Error())
	}

	return app.ResponseMsg("添加成功")
}
