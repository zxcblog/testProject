package api

import (
	"bytes"
	"net/http"
	"time"

	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12"
)

type CaptchaController struct {
	Ctx iris.Context
}

// Post 获取验证码图片步骤2
// @Summary 获取验证码图片
// @Description 前台获取验证码
// @Accept json
// @Produce json
// @Tags 验证码
// @Success 200
// @Router /api/captcha/{captcha_id} [get]
func (c *CaptchaController) GetBy(captcha_id string) {

	if c.Ctx.URLParam("t") != "" {
		captcha.Reload(captcha_id)
	}
	c.responseCaptchaImage(captcha_id, 200, 50)
}

func (c *CaptchaController) responseCaptchaImage(id string, width, height int) error {
	w := c.Ctx.ResponseWriter()
	r := c.Ctx.Request()

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	ext := ".png"
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, "zh")
	default:
		return captcha.ErrNotFound
	}

	download := false
	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
