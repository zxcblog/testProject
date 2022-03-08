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

type PostRegisterCheckRequest struct {
	UserName        string `validate:"required,min=4,max=32" label:"账号" json:"userName"`                         // 账号
	Nickname        string `validate:"required,min=1,max=16" label:"昵称" json:"nikeName"`                         // 昵称
	Password        string `validate:"required,min=4,max=16" label:"昵称" json:"passWord"`                         // 密码
	ConfirmPassword string `validate:"required,min=4,max=16,eqfield=Password" label:"昵称" json:"confirmPassWord"` // 确认密码
}

type PostLoginCheckRequest struct {
	UserName string `validate:"required,min=4,max=32" label:"账号" json:"userName"` // 账号
	Password string `validate:"required,min=4,max=16" label:"昵称" json:"passWord"` // 密码
}

// Post 用户注册
// @Summary 用户注册
// @Description 前台用户注册
// @Accept json
// @Produce json
// @param root body PostRegisterCheckRequest true "用户注册"
// @Tags 用户注册
// @Success 200 {object} app.Response "注册成功"
// @Router /api/user/register [post]
func (u *UserController) PostRegister() *app.Response {
	params := &PostRegisterCheckRequest{}

	//绑定参数
	if err := u.Ctx.ReadJSON(params); err != nil {
		return app.ResponseErrMsg(err.Error())
	}

	//参数校验
	if err := global.Validate.ValidateParam(params); err != nil {
		return app.ToResponseErr(errcode.InvalidParams.SetMsg(err.Error()))
	}

	//绑定model参数
	user := &models.User{
		Username: params.UserName,
		Nickname: params.Nickname,
		Password: params.Password,
	}

	//调用services逻辑处理层
	if err := services.UserService.Create(user); err != nil {
		return app.ToResponseErr(err)
	}

	//返回用户登录信息
	//先写登录 直接调用登录方法
	return app.ResponseMsg("注册成功")
}

//用户登录
func (u *UserController) PostLogin() *app.Response {
	params := &PostLoginCheckRequest{}

	//绑定参数

	if err := u.Ctx.ReadJSON(params); err != nil {
		return app.ResponseErrMsg(err.Error())
	}

	//校验账户密码

	return app.ResponseErrMsg("111")
}
