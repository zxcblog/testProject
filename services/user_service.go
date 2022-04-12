package services

import (
	"errors"
	"new-project/cache"
	"new-project/global"
	"new-project/models"
	"new-project/pkg/util"
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

// UniqueByName 查看用户名唯一， 返回true时代表当前账号已被注册
func (u *userService) UniqueByName(username string) bool {
	return repositories.UserRepositories.GetUsernameData(global.DB, username) != nil
}

// Create 创建用户信息
func (u *userService) Create(user *models.User) error {
	if u.UniqueByName(user.Username) {
		return errors.New("当前账号已被注册")
	}

	//随机数
	user.Salt = util.RandomStr(16)
	//密码加盐
	user.Password = util.Md5Encrypt(user.Password + user.Salt)

	if err := repositories.UserRepositories.Create(global.DB, user); err != nil {
		global.Logger.Error("注册失败", zap.Error(err))
		return errors.New("用户注册失败")
	}

	// 通过账号(唯一)将用户放到缓存中
	cache.UserCache.Set(user)
	return nil
}

//Login 用户账号密码登录
func (u *userService) Login(username, password string) (*models.User, error) {
	user := repositories.UserRepositories.GetUsernameData(global.DB, username)
	if user == nil {
		return nil, errors.New("用户名或密码错误")
	}

	//校验密码
	if util.Md5Encrypt(password+user.Salt) != user.Password {
		return nil, errors.New("用户名或密码错误")
	}
	return user, nil
}
