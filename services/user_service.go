package services

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"new-project/global"
	"new-project/models"
	"new-project/pkg/errcode"
	"new-project/repositories"
	"strings"
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

func (u *userService) Create(user *models.User) error {
	//随机数
	user.Salt = getRandstring(16)
	//密码加盐
	user.Password = MD5_SALT(user.Password, user.Salt)
	err := repositories.UserRepositories.Create(global.DB, user)
	if err != nil {
		global.Logger.Error("注册失败", zap.Error(err))
		return errcode.CreateError.SetMsg("注册失败")
	}
	return nil
}

//md5密码加盐 到时候建一个全局function这些方法放在那里
func MD5_SALT(str string, salt string) string {
	b := []byte(str)
	s := []byte(salt)
	h := md5.New()
	h.Write(s) // 先写盐值
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

//取得随机字符串:使用字符串拼接
func getRandstring(length int) string {
	if length < 1 {
		return ""
	}
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charArr := strings.Split(char, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))

	var rchar string = ""
	for i := 1; i <= length; i++ {
		rchar = rchar + charArr[ran.Intn(charlen)]
	}
	return rchar
}
