package api

import (
	"new-project/controller/render"
	"new-project/global"
	"new-project/models"
	"new-project/models/form"
	"new-project/pkg/app"
	"new-project/services"

	"github.com/kataras/iris/v12"
)

type UserController struct {
	Ctx iris.Context
}

// PostUniqueBy  验证用户名唯一
// @Summary      验证用户名唯一
// @Description  验证用户名唯一
// @Accept       json
// @Produce      json
// @param        userName  path  string  true  "账号"
// @Tags         用户
// @Success      200  {object}  app.Response{data=bool}
// @Router       /user/unique/{userName} [post]
func (this *UserController) PostUniqueBy(userName string) *app.Response {
	return app.ResponseData(services.UserService.UniqueByName(userName))
}

// PostRegister  用户注册
// @Summary      用户注册
// @Description  前台用户注册
// @Accept       json
// @Produce      json
// @param        root  body  form.UserRegister  true  "用户注册"
// @Tags         用户
// @Success      200  {object}  app.Response{data=app.Result{token=services.Token,user=render.User}}
// @Router       /user/register [post]
func (this *UserController) PostRegister() *app.Response {
	params := &form.UserRegister{}
	if err := app.FormValueJson(this.Ctx, global.Validate, params); err != nil {
		return app.ResponseErrMsg(err.Error())
	}

	// 验证码校验
	if flag, err := services.CaptchaService.VerifyCaptcha(this.Ctx, params.CaptchaId, params.UserCapt); err != nil || !flag {
		return app.ResponseErrMsg("验证码错误")
	}

	//绑定model参数
	user := &models.User{
		Avatar:   params.Avatar,
		Username: params.UserName,
		Nickname: params.Nickname,
		Password: params.Password,
	}
	if err := services.UserService.Create(user); err != nil {
		return app.ResponseErrMsg("用户创建失败")
	}
	return render.BuildLoginSuccess(user)
}

type PostLoginCheckRequest struct {
	UserName  string `validate:"required,min=4,max=32" label:"账号" json:"userName"` // 账号
	Password  string `validate:"required,min=4,max=16" label:"昵称" json:"passWord"` // 密码
	CaptchaId string `validate:"required" label:"验证码id" json:"captchaId"`          // 验证码id
	UserCapt  string `validate:"required" label:"验证码" json:"userCapt"`             // 验证码
}

// PostLogin 用户登录
// @Summary      用户登录
// @Description  前台用户登录
// @Accept       json
// @Produce      json
// @param        root  body  PostLoginCheckRequest  true  "用户登录"
// @Tags         用户
// @Success      200  {object}  app.Response{data=app.Result}
// @Router       /user/login [post]
func (this *UserController) PostLogin() *app.Response {
	params := &PostLoginCheckRequest{}
	if err := app.FormValueJson(this.Ctx, global.Validate, params); err != nil {
		return app.ToResponseErr(err)
	}

	if flag, err := services.CaptchaService.VerifyCaptcha(this.Ctx, params.CaptchaId, params.UserCapt); err != nil || !flag {
		return app.ResponseErrMsg("验证码错误")
	}

	user, err := services.UserService.Login(params.UserName, params.Password)
	if err != nil {
		return app.ToResponseErr(err)
	}

	token, err := services.UserService.GenTokenDefault2Hour(user)
	if err != nil {
		return app.ToResponseErr(err)
	}
	return app.ToResponse("注册成功", app.Result{
		"token": token,
		"user":  render.BuildUser(user),
	})
}
