// Package models
// @Author:        asus
// @Description:   $
// @File:          shop
// @Data:          2022/2/2613:07
//
package models

var Models = []interface{}{
	&Product{}, &User{}, &Category{},
}

// User 用户表
type User struct {
	Model
	Status   uint   `gorm:"type:tinyint(1);index:idx_shop_user_status;not null;comment:状态"`
	Username string `gorm:"size:32;unique;comment:账号"`
	Nickname string `gorm:"size:16;comment:昵称"`
	Avatar   string `gorm:"type:text;comment:头像"`
	Password string `gorm:"size:512;comment:密码"`
	Salt     string `gorm:"size:16;comment:密码加盐"`
}

// Category 商品分类表
type Category struct {
	Model
	Path         string `gorm:"size:512;default:-;comment:路径"`
	CategoryName string `gorm:"size:512;index:category_name_and_is_final;comment:分类名称"`
	CategoryImg  string `gorm:"type:text;comment:分类图片"`
	IsFinal      bool   `gorm:"default:false;index:category_name_and_is_final;comment:是否为终极类目"`
	Sort         uint   `gorm:"default:0;comment:排序"`
	CategoryID   uint   `gorm:"index:idx_category_shop_category_id;default:0;comment:父级类目"`
	Level        uint   `gorm:"default:1;comment:级别"`
}

// Product 商品表
type Product struct {
	Model
}
