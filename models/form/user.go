// Package form
// @Author:        asus
// @Description:   $
// @File:          user
// @Data:          2022/4/1116:57
//
package form

// UserRegister 用户注册
type UserRegister struct {
	UserName        string `validate:"required,min=6,max=32" label:"账号" json:"userName"`                           // 账号
	Nickname        string `validate:"required,min=1,max=16" label:"昵称" json:"nikeName"`                           // 昵称
	Password        string `validate:"required,min=6,max=16" label:"密码" json:"passWord"`                           // 密码
	ConfirmPassword string `validate:"required,min=6,max=16,eqfield=Password" label:"确认密码" json:"confirmPassWord"` // 确认密码
	CaptchaId       string `validate:"required" label:"验证码id" json:"captchaId"`                                    // 验证码id
	UserCapt        string `validate:"required" label:"验证码" json:"userCapt"`                                       // 验证码
}

// UserLogin 用户登录
type UserLogin struct {
	UserName  string `validate:"required,min=4,max=32" label:"账号" json:"userName"` // 账号
	Password  string `validate:"required,min=4,max=16" label:"昵称" json:"passWord"` // 密码
	CaptchaId string `validate:"required" label:"验证码id" json:"captchaId"`          // 验证码id
	UserCapt  string `validate:"required" label:"验证码" json:"userCapt"`             // 验证码
}
