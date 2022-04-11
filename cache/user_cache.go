package cache

import (
	"context"
	"fmt"
	"new-project/global"
	"new-project/models"
	"new-project/pkg/util"
	"time"

	"go.uber.org/zap"
)

var UserCache = NewUserCache(context.Background(), "user:")

type userCache struct {
	ctx context.Context
	key string
}

func NewUserCache(ctx context.Context, key string) *userCache {
	return &userCache{ctx: ctx, key: key}
}

//设置用户登录缓存信息
func (this *userCache) SetUserLoginData(token string, user *models.User) (bool, error) {
	field := this.key + token

	// 添加缓存
	// 设置redis过期时间为2个半小时， 用户在过期半小时以内使用当前token更换新token, 不用重新登录
	res, err := global.Redis.Set(this.ctx, field, util.StructToString(user), 2*time.Hour+3.*time.Second).Result()
	if err != nil {
		global.Logger.Error("token缓存用户信息失败：", zap.Error(err))
		return false, err
	}

	return res != "", nil
}

func (this *userCache) Set(user *models.User) {
	res, err := global.Redis.Set(this.ctx, this.key+user.Username, util.StructToString(user), 2*time.Hour+30*time.Second).Result()
	fmt.Println(res)
	if err != nil {
		global.Logger.Error("[cache] 用户信息存储失败")
	}
}

func (this *userCache) Get(username string) (user *models.User) {
	res, err := global.Redis.Get(this.ctx, this.key+username).Result()
	if err != nil {
		global.Logger.Error("[cache] 用户信息读取失败")
		return nil
	}

	util.StringToStruct(res, &user)
	return user
}
