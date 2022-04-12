// Package form
// @Author:        asus
// @Description:   $
// @File:          form
// @Data:          2022/4/1210:15
//
package form

// Brand 品牌管理修改创建参数
type Brand struct {
	BrandName   string `validate:"required,min=1,max=20" label:"品牌名称" json:"brandName"` // 品牌名称
	BrandImg    string `validate:"-" label:"品牌图片" json:"brandImg" default:""`           // 品牌图片地址链接
	Description string `validate:"-" label:"品牌描述" json:"description" default:""`        // 品牌描述
	CategoryID  uint   `validate:"required" label:"所属分类" json:"categoryID"`             // 所属分类
}

// Category 分类管理修改创建参数
type Category struct {
	CategoryName string `validate:"required,min=1,max=20" label:"分类名称" json:"categoryName"` // 分类名称
	CategoryImg  string `validate:"-" label:"分类图片" json:"categoryImg" default:""`           // 分类图片地址链接
	CategoryID   uint   `validate:"-" label:"所属分类" json:"categoryID" default:"0"`           // 所属分类
	Sort         uint   `validate:"min=0,max=100" label:"排序" json:"sort" default:"0"`       //排序
}

// CategoryQueryName 通过名称查询分类列表
type CategoryQueryName struct {
	CategoryName string `validate:"required,min=1,max=20" label:"分类名称" json:"categoryName"` // 分类名称
}
