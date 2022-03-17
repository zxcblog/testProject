package util

import (
	"github.com/dchest/captcha"
)

type MyCaptchaModel struct {
	Id   string
	Path string
}

// 几位数字的验证码
func Captcha(len int) *MyCaptchaModel {
	captchaId := captcha.NewLen(len)
	stru := MyCaptchaModel{
		Id:   captchaId,
		Path: "/api/captcha/" + captchaId + "",
	}
	return &stru
}

func VerifyCaptcha(captchaId, postCaptcha string) bool {
	return captcha.VerifyString(captchaId, postCaptcha)
}
