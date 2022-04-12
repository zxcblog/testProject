// Package cache
// @Author:        asus
// @Description:   $
// @File:          captchCache
// @Data:          2022/4/129:39
//
package cache

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"new-project/global"
	"time"
)

var CaptchCache = NewCaptchCache(context.Background(), "captcha:")

type captchCache struct {
	ctx context.Context
	key string
}

func NewCaptchCache(ctx context.Context, key string) *captchCache {
	return &captchCache{ctx: ctx, key: key}
}

func (this *captchCache) Set(randId string, value []byte) {
	res, err := global.Redis.Set(this.ctx, this.key+randId, value, 2*time.Hour).Result()
	fmt.Println(res)
	if err != nil {
		global.Logger.Error("[cache] 缓存验证码id失败", zap.Error(err))
	}
}

// IsExists 判断验证码唯一id是否存在
func (this *captchCache) IsExists(randId string) bool {
	flag := global.Redis.Exists(this.ctx, this.key+randId).Val()
	return flag > 0
}

func (this *captchCache) Get(randId string) string {
	str, err := global.Redis.Get(this.ctx, this.key+randId).Result()
	if err != nil {
		global.Logger.Error("[cache] 验证码缓存获取失败", zap.Error(err))
		return ""
	}

	return str
}

func (this *captchCache) Del(captchaId string) {
	global.Redis.Del(this.ctx, this.key+captchaId)
}
