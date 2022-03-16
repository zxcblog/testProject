package api

import (
	"new-project/global"
	"new-project/models"
	"new-project/pkg/app"
	"new-project/pkg/util"
	"new-project/services"

	"github.com/kataras/iris/v12"
)

type UserController struct {
	Ctx iris.Context
}

type JwtResponseData struct {
	Toekn string `json:"token"`
}

type CaptchResponseData struct {
	CaptchaPath string `json:"captchaPath"`
	CaptchaId   string `json:"captchaId"`
}

type PostRegisterCheckRequest struct {
	UserName        string `validate:"required,min=4,max=32" label:"账号" json:"userName"`                         // 账号
	Nickname        string `validate:"required,min=1,max=16" label:"昵称" json:"nikeName"`                         // 昵称
	Password        string `validate:"required,min=4,max=16" label:"昵称" json:"passWord"`                         // 密码
	ConfirmPassword string `validate:"required,min=4,max=16,eqfield=Password" label:"昵称" json:"confirmPassWord"` // 确认密码
	CaptchaId       string `validate:"required" label:"验证码id" json:"captchaId"`                                  // 验证码id
	UserCapt        string `validate:"required" label:"验证码" json:"userCapt"`                                     // 验证码
}

type PostLoginCheckRequest struct {
	UserName  string `validate:"required,min=4,max=32" label:"账号" json:"userName"` // 账号
	Password  string `validate:"required,min=4,max=16" label:"昵称" json:"passWord"` // 密码
	CaptchaId string `validate:"required" label:"验证码id" json:"captchaId"`          // 验证码id
	UserCapt  string `validate:"required" label:"验证码" json:"userCapt"`             // 验证码
}

// Post 用户注册
// @Summary 用户注册
// @Description 前台用户注册
// @Accept json
// @Produce json
// @param root body PostRegisterCheckRequest true "用户注册"
// @Tags 用户
// @Success 200 {object} app.Response{data=JwtResponseData}
// @Router /api/user/register [post]
func (this *UserController) PostRegister() *app.Response {
	params := &PostRegisterCheckRequest{}

	//参数校验 && 绑定参数
	if err := app.FormValueJson(this.Ctx, global.Validate, params); err != nil {
		return app.ToResponseErr(err)
	}

	if !util.VerifyCaptcha(params.CaptchaId, params.UserCapt) {
		return app.ResponseErrMsg("验证码错误")
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
	// 生成Token
	tokenString, _ := app.GenToken(user)

	return app.ToResponse("注册成功", JwtResponseData{
		Toekn: tokenString,
	})
}

// Post 获取验证码步骤1
// @Summary 获取验证码信息
// @Description 前台获取验证码信息
// @Accept json
// @Produce json
// @Tags 验证码
// @Success 200 {object} app.Response{data=CaptchResponseData}
// @Router /api/user/login [get]
func (this *UserController) GetLogin() *app.Response {
	capt := util.Captcha(4)
	return app.ToResponse("获取成功", CaptchResponseData{
		CaptchaPath: capt.Path,
		CaptchaId:   capt.Id,
	})
}

// Post 用户登录
// @Summary 用户登录
// @Description 前台用户登录
// @Accept json
// @Produce json
// @param root body PostLoginCheckRequest true "用户登录"
// @Tags 用户
// @Success 200 {object} app.Response{data=JwtResponseData}
// @Router /api/user/login [post]
func (this *UserController) PostLogin() *app.Response {
	params := &PostLoginCheckRequest{}

	if err := app.FormValueJson(this.Ctx, global.Validate, params); err != nil {
		return app.ToResponseErr(err)
	}

	if !util.VerifyCaptcha(params.CaptchaId, params.UserCapt) {
		return app.ResponseErrMsg("验证码错误")
	}

	user := &models.User{
		Username: params.UserName,
		Password: params.Password,
	}

	//调用services逻辑层  校验账户密码
	if err := services.UserService.Login(user); err != nil {
		return app.ToResponseErr(err)
	}

	// 生成Token
	tokenString, err := app.GenToken(user)
	if err != nil {
		return app.ToResponseErr(err)
	}

	//返回参数
	return app.ToResponse("登录成功", JwtResponseData{
		Toekn: tokenString,
	})
}
