package api

import (
	"bytes"
	"net/http"
	"new-project/pkg/app"
	"new-project/services"
	"time"

	"github.com/kataras/iris/v12"
)

type CaptchaController struct {
	Ctx iris.Context
}

// Get 获取验证码id
// @Summary      获取验证码id
// @Description  获取验证码id
// @Produce      json
// @Tags         验证码
// @Success      200  {object}  app.Response{data=string}
// @Router       /api/captcha [get]
func (this *CaptchaController) Get() *app.Response {
	return app.ResponseData(services.CaptchaService.GetCaptchaID(this.Ctx))
}

// GetBy 通过验证码ID获取到图片
// @Summary      通过验证码ID获取到图片
// @Description  通过验证码ID获取到图片
// @param        refresh  query  bool  false  "是否刷新验证码"
// @Produce      json
// @Tags         验证码
// @Success      200
// @Router       /api/captcha/{captchaId} [get]
func (this *CaptchaController) GetBy(captchaId string) *app.Response {
	refresh, _ := this.Ctx.URLParamBool("refresh")
	var content bytes.Buffer
	err := services.CaptchaService.GetImage(this.Ctx, &content, captchaId, refresh, 200, 50)
	if err != nil {
		return app.ToResponseErr(err)
	}

	w := this.Ctx.ResponseWriter()
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "image/png")
	http.ServeContent(w, this.Ctx.Request(), captchaId+".png", time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
