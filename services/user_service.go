package services

import (
	"new-project/cache"
	"new-project/global"
	"new-project/models"
	"new-project/pkg/app"
	"new-project/pkg/encrypt"
	"new-project/pkg/errcode"
	"new-project/pkg/util"
	"new-project/repositories"
	"time"

	"go.uber.org/zap"
)

var UserService = NewUserService()

//定义user结构体
type userService struct{}

//相当于工厂模式
func NewUserService() *userService {
	return &userService{}
}

// UniqueByName 查看用户名唯一， 返回true时代表当前账号已被注册
func (u *userService) UniqueByName(username string) bool {
	return repositories.UserRepositories.GetUsernameData(global.DB, username) != nil
}

// Create 创建用户信息
func (u *userService) Create(user *models.User) error {
	if u.UniqueByName(user.Username) {
		return errcode.CreateError.SetMsg("当前账号已被注册")
	}

	//随机数
	user.Salt = util.RandomStr(16)
	//密码加盐
	user.Password = encrypt.Md5Encrypt(user.Password + user.Salt)
	err := repositories.UserRepositories.Create(global.DB, user)
	if err != nil {
		global.Logger.Error("注册失败", zap.Error(err))
	}

	return err
}

type Token struct {
	Token    string `json:"token"`
	Expire   int64  `json:"expire"`
	Duration int64  `json:"duration"`
}

// GenToken 生成token并缓存用户信息
func (u *userService) GenTokenDefault2Hour(user *models.User) (*Token, error) {
	//生成Token
	tokenExpire := time.Now().Add(2 * time.Hour)
	tokenString, err := app.GenToken(user, tokenExpire)
	if err != nil {
		global.Logger.Error("用户注册生成token失败", zap.Error(err))
		return nil, errcode.UnauthorizedTokenGenerate
	}

	//将用户信息储存到redis
	if _, err := cache.UserCache.SetUserLoginData(tokenString, user); err != nil {
		global.Logger.Error("用户信息缓存token失败", zap.Error(err))
		return nil, errcode.UnauthorizedTokenGenerate
	}

	return &Token{
		Token:    tokenString,
		Expire:   tokenExpire.Unix(),
		Duration: 2 * time.Hour.Milliseconds() / 1000,
	}, nil
}

//调用model 验证账号和密码 调用jwt
func (u *userService) Login(username, password string) (*models.User, error) {
	user := repositories.UserRepositories.GetUsernameData(global.DB, username)
	if user == nil {
		return nil, errcode.CreateError.SetMsg("用户名或密码错误")
	}

	//校验密码
	if encrypt.Md5Encrypt(password+user.Salt) != user.Password {
		return nil, errcode.CreateError.SetMsg("用户名或密码错误")
	}

	return user, nil
}
