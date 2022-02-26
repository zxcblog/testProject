// Package models
// @Author:        asus
// @Description:   $
// @File:          shop
// @Data:          2022/2/2613:07
//
package models

import "database/sql"

var Models = []interface{}{
	&Product{}, &User{},
}

type Product struct {
	Model
}

type User struct {
	Model
	Status   uint           `gorm:"type:tinyint(1);index:idx_shop_user_status;not null"` // 状态
	Username sql.NullString `gorm:"size:32;unique"`                                      // 账号
	Nickname string         `gorm:"size:16"`                                             // 昵称
	Avatar   string         `gorm:"type:text"`                                           // 头像
	Password string         `gorm:"size:512"`                                            // 密码
	Salt     string         `gorm:"size:16"`                                             // 密码加盐
}
