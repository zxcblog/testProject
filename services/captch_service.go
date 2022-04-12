package services

import (
	"errors"
	"github.com/dchest/captcha"
	"go.uber.org/zap"
	"io"
	"new-project/cache"
	"new-project/global"
	"new-project/pkg/util"
	"strconv"
	"time"
)

var CaptchaService = NewCaptchaService()

type captchaService struct{}

func NewCaptchaService() *captchaService {
	return &captchaService{}
}

// GetCaptchaID 获取验证码ID
func (this *captchaService) GetCaptchaID() string {
	// 获取随机字符串值
	randID := time.Now().Format("150405") + util.RandomStr(14)

	cache.CaptchCache.Set(randID, captcha.RandomDigits(4))
	return randID
}

// GetImage 获取验证码图片
func (this *captchaService) GetImage(w io.Writer, captchaId string, refresh bool, width, height int) error {
	if refresh {
		if cache.CaptchCache.IsExists(captchaId) {
			cache.CaptchCache.Set(captchaId, captcha.RandomDigits(4))
		} else {
			return errors.New("验证码不存在，请重新获取")
		}
	}

	val := cache.CaptchCache.Get(captchaId)
	if val == "" {
		return errors.New("验证码不存在，请重新获取")
	}

	_, err := captcha.NewImage(captchaId, []byte(val), width, height).WriteTo(w)
	if err != nil {
		global.Logger.Error("[service] 验证码转图片失败", zap.Error(err))
		return errors.New("验证码不存在，请重新获取")
	}
	return nil
}

// VerifyCaptcha 验证码校验
func (this *captchaService) VerifyCaptcha(captchaId, userCaptcha string) error {
	val := cache.CaptchCache.Get(captchaId)
	if val == "" {
		return errors.New("验证码不存在，请重新获取")
	}

	for i, c := range val {
		if string(userCaptcha[i]) != strconv.Itoa(int(c)) {
			return errors.New("验证码不正确")
		}
	}

	cache.CaptchCache.Del(captchaId)
	return nil
}
