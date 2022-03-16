package app

import (
	"errors"
	"new-project/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	Username string      `json:"username"`
	UserInfo models.User `json:"userInfo"`
	jwt.StandardClaims
}

//设置过期时间
const TokenExpireDuration = time.Hour * 2

//设置本地密钥
var MySecret = []byte("j6cFIiBsY9jSIxXg")

// GenToken 生成JWT
func GenToken(user *models.User) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		user.Username, // 自定义字段
		*user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),                          //发放时间
			Issuer:    "test-project",                             // 签发人
		},
	}
	// 使用指定的签名方法(hash)创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
