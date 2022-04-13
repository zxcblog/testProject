package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"new-project/cache"
	"new-project/models"
	"new-project/pkg/app"
	"new-project/pkg/config"
	"time"
)

var UserTokenService = NewUserTokenService()

type userTokenService struct{}

func NewUserTokenService() *userTokenService {
	return &userTokenService{}
}

type LoginClaims struct {
	UserName string
	jwt.StandardClaims
}

// TokenGenerate 生成token
func (this *userTokenService) TokenGenerate(userName string, issuedAt time.Time) (string, error) {
	expire := issuedAt.Add(config.GetJwtExpire())
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, LoginClaims{
		userName,
		jwt.StandardClaims{
			ExpiresAt: expire.Unix(),               // 过期时间
			IssuedAt:  issuedAt.Unix(),             //发放时间
			Issuer:    config.GetService().AppName, // 签发人
		},
	})

	return token.SignedString(config.GetJwtSecret())
}

// ParseToken 解析Token，获取登录jwt结构体
func (*userTokenService) ParseToken(tokenString string) (*LoginClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.GetJwtSecret(), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*LoginClaims)
	if !ok && !token.Valid {
		return nil, errors.New("token失败")
	}

	return claims, nil
}

// GetTokenUserName 解析Token, 获取token中的用户名称
func (this *userTokenService) GetTokenUserName(token string) (string, error) {
	claims, err := this.ParseToken(token)
	if err != nil {
		return "", err
	}
	return claims.UserName, nil
}

// GetTokenUser 解析token，获取到当前登录用户信息
func (this *userTokenService) GetTokenUser(token string) (*models.User, *app.Response) {
	claims, err := this.ParseToken(token)
	if err != nil {
		return nil, app.UnauthorizedTokenError
	}

	user := cache.UserCache.Get(claims.UserName)
	if user == nil {
		return nil, app.UnauthorizedAuthNotExist
	}
	return user, nil
}
