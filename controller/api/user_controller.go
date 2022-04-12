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
// @Success      200  {object}  app.Response{data=app.Result{user=render.User}}
// @Router       /user/register [post]
func (this *UserController) PostRegister() *app.Response {
	params := &form.UserRegister{}
	if err := app.FormValueJson(this.Ctx, global.Validate, params); err != nil {
		return err
	}

	// 验证码校验
	if err := services.CaptchaService.VerifyCaptcha(params.CaptchaId, params.UserCapt); err != nil {
		return app.ResponseErrMsg(err.Error())
	}

	//绑定model参数
	user := &models.User{
		Username: params.UserName,
		Nickname: params.Nickname,
		Password: params.Password,
	}
	if err := services.UserService.Create(user); err != nil {
		return app.CreateError.SetMsg(err.Error())
	}
	return render.BuildLoginSuccess(user)
}

// PostLogin 用户登录
// @Summary      用户登录
// @Description  前台用户登录
// @Accept       json
// @Produce      json
// @param        root  body  form.UserLogin  true  "用户登录"
// @Tags         用户
// @Success      200  {object}  app.Response{data=app.Result}
// @Router       /user/login [post]
func (this *UserController) PostLogin() *app.Response {
	params := &form.UserLogin{}
	if err := app.FormValueJson(this.Ctx, global.Validate, params); err != nil {
		return err
	}

	if err := services.CaptchaService.VerifyCaptcha(params.CaptchaId, params.UserCapt); err != nil {
		return app.ResponseErrMsg(err.Error())
	}

	user, err := services.UserService.Login(params.UserName, params.Password)
	if err != nil {
		return app.ResponseErrMsg(err.Error())
	}
	return render.BuildLoginSuccess(user)
}
