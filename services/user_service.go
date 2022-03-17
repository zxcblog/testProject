package services

import (
	"new-project/cache"
	"new-project/global"
	"new-project/models"
	"new-project/pkg/app"
	"new-project/pkg/errcode"
	"new-project/repositories"

	"go.uber.org/zap"
)

var UserService = NewUserService()

//定义user结构体
type userService struct{}

//相当于工厂模式
func NewUserService() *userService {
	return &userService{}
}

func (u *userService) Create(user *models.User) (string, error) {
	//随机数
	user.Salt = app.GetRandString(16)
	//密码加盐
	user.Password = app.Md5Salt(user.Password, user.Salt)
	err := repositories.UserRepositories.Create(global.DB, user)
	if err != nil {
		global.Logger.Error("注册失败", zap.Error(err))
		return "", errcode.CreateError.SetMsg("注册失败")
	}

	//随机过期时间
	tokenExpireDuration := app.CacheTimeGenerator(120, 30)

	// 生成Token
	tokenString, err := app.GenToken(user, tokenExpireDuration)
	if err != nil {
		return "", err
	}
	//将用户信息储存到redis
	cache.UserCache.SetUserLoginData(tokenString, tokenExpireDuration, user)

	return tokenString, nil
}

//调用model 验证账号和密码 调用jwt
func (u *userService) Login(user *models.User) (string, error) {
	userData := repositories.UserRepositories.GetUsernameData(global.DB, user.Username)
	if userData == nil {
		return "", errcode.CreateError.SetMsg("用户名或密码错误")
	}

	//校验密码
	if userCheckPassWord := app.Md5Salt(user.Password, userData.Salt); userCheckPassWord != userData.Password {
		return "", errcode.CreateError.SetMsg("密码错误")
	}

	if userData.Status == 2 {
		return "", errcode.CreateError.SetMsg("账号状态异常")
	}
	//随机过期时间
	tokenExpireDuration := app.CacheTimeGenerator(120, 30)

	// 生成Token
	tokenString, err := app.GenToken(userData, tokenExpireDuration)
	if err != nil {
		return "", err
	}
	//将用户信息储存到redis
	cache.UserCache.SetUserLoginData(tokenString, tokenExpireDuration, userData)

	return tokenString, nil
}
