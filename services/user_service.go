package services

import (
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

func (u *userService) Create(user *models.User) error {
	//随机数
	user.Salt = app.GetRandString(16)
	//密码加盐
	user.Password = app.Md5Salt(user.Password, user.Salt)
	err := repositories.UserRepositories.Create(global.DB, user)
	if err != nil {
		global.Logger.Error("注册失败", zap.Error(err))
		return errcode.CreateError.SetMsg("注册失败")
	}
	return nil
}
