package api

import (
	"new-project/global"
	"new-project/models"
	"new-project/pkg/app"
	"new-project/pkg/errcode"
	"new-project/services"

	"github.com/kataras/iris/v12"
)

type UserController struct {
	Ctx iris.Context
}

type PostRegisterCkeckRequest struct {
	UserName        string `validate:"required,min=4,max=32" label:"账号" json:"userName"`        // 账号
	Nickname        string `validate:"required,min=1,max=16" label:"昵称" json:"nikeName"`        // 昵称
	Password        string `validate:"required,min=4,max=16" label:"昵称" json:"passWord"`        // 密码
	ConfirmPassword string `validate:"required,min=4,max=16" label:"昵称" json:"confirmPassWord"` // 确认密码
}

// Post 用户注册
// @Summary 用户注册
// @Description 前台用户注册
// @Accept json
// @Produce json
// @param root body PostRegisterCkeckRequest true "用户注册"
// @Tags 商品分类
// @Success 200 {object} app.Response "注册成功"
// @Router /api/user/register [post]
func (u *UserController) PostRegister() *app.Response {
	parmas := &PostRegisterCkeckRequest{}

	//绑定参数
	if err := u.Ctx.ReadJSON(parmas); err != nil {
		return app.ResponseErrMsg(err.Error())
	}

	//参数校验
	if err := global.Validate.ValidateParam(parmas); err != nil {
		return app.ToResponseErr(errcode.InvalidParams.SetMsg(err.Error()))
	}

	//校验两次密码是否一致
	if parmas.Password != parmas.ConfirmPassword {
		return app.ResponseErrMsg("密码与确认密码不一致")
	}

	//绑定model参数
	user := &models.User{
		Username: parmas.UserName,
		Nickname: parmas.Nickname,
		Password: parmas.Password,
	}

	//调用services逻辑处理层
	if err := services.UserService.Create(user); err != nil {
		return app.ToResponseErr(err)
	}
	return app.ResponseMsg("注册成功")
}
