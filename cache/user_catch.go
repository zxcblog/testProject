package cache

import (
	"context"
	"new-project/global"
	"new-project/models"
	"new-project/pkg/util"
	"time"

	"go.uber.org/zap"
)

var UserCache = NewUserCache(context.Background(), "user:login:")

type userCatch struct {
	ctx context.Context
	key string
}

func NewUserCache(ctx context.Context, key string) *userCatch {
	return &userCatch{ctx: ctx, key: key}
}

//设置用户登录缓存信息
func (this *userCatch) SetUserLoginData(token string, tokenExpireDuration time.Duration, user *models.User) (bool, error) {
	field := this.key + token

	// 添加缓存
	res, err := global.Redis.Set(this.ctx, field, util.StructToString(user), tokenExpireDuration).Result()
	if err != nil {
		global.Logger.Error("分类缓存失败：", zap.Error(err))
		return false, err
	}

	return res != "", nil
}
