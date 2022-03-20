package services

import (
	"context"
	"github.com/dchest/captcha"
	"go.uber.org/zap"
	"io"
	"new-project/global"
	"new-project/pkg/errcode"
	"new-project/pkg/util"
	"strconv"
	"time"
)

const cacheKey string = "captcha:"

var CaptchaService = NewCaptchaService()

type captchaService struct{}

func NewCaptchaService() *captchaService {
	return &captchaService{}
}

func (*captchaService) refresh(ctx context.Context, randID string) {
	// 将 验证码id 保存5分钟， 用户超过5分钟没有对验证码进行验证， 则验证码错误
	global.Redis.Set(ctx, cacheKey+randID, captcha.RandomDigits(4), 5*time.Minute)
}

// GetCaptchaID 获取验证码ID
func (this *captchaService) GetCaptchaID(ctx context.Context) string {
	// 获取随机字符串值
	randID := time.Now().Format("150405") + util.RandomStr(14)

	this.refresh(ctx, randID)
	return randID
}

// GetImage 获取验证码图片
func (this *captchaService) GetImage(ctx context.Context, w io.Writer, captchaId string, refresh bool, width, height int) error {
	key := cacheKey + captchaId
	isExists, _ := global.Redis.Exists(ctx, key).Result()
	if isExists != 1 {
		return errcode.NotFound.SetMsg("验证码不存在，请重新获取")
	}

	// 刷新验证码
	if refresh {
		this.refresh(ctx, captchaId)
	}

	// 获取验证码值
	val, err := global.Redis.Get(ctx, key).Result()
	if err != nil || val == "" {
		return errcode.NotFound.SetMsg("验证码不存在，请重新获取")
	}

	_, err = captcha.NewImage(captchaId, []byte(val), width, height).WriteTo(w)
	if err != nil {
		global.Logger.Error("验证码转图片失败", zap.Error(err))
		return errcode.NotFound.SetMsg("验证码不存在，请重新获取")
	}

	return nil
}

//captcha.VerifyString(captchaId, postCaptcha)

// VerifyCaptcha 验证码校验
func (this *captchaService) VerifyCaptcha(ctx context.Context, captchaID, userCaptcha string) (bool, error) {
	// 获取验证码值
	val, err := global.Redis.Get(ctx, cacheKey+captchaID).Result()
	if err != nil || val == "" {
		return false, errcode.NotFound.SetMsg("验证码已过期，请重新编辑")
	}

	for i, c := range val {
		if string(userCaptcha[i]) != strconv.Itoa(int(c)) {
			// 验证码校验错误， 刷新验证码，让用户重新校验
			this.refresh(ctx, captchaID)
			return false, errcode.NotFound.SetMsg("输入验证码不正确")
		}
	}

	global.Redis.Del(ctx, cacheKey+captchaID)
	return true, nil
}
