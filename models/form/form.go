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

// Address 用户地址创建修改
type Address struct {
	IsDefault     bool   `validate:"-" label:"默认地址" json:"isDefault"`                    // 默认地址
	Province      string `validate:"required" label:"省" json:"province"`                 // 省
	City          string `validate:"required" label:"市" json:"city"`                     // 市
	Area          string `validate:"required" label:"区/县" json:"area"`                   // 区/县
	Street        string `validate:"-" label:"街道" json:"street"`                         // 街道
	Desc          string `validate:"required" label:"详细地址" json:"desc"`                  // 详细地址
	ContactName   string `validate:"required" label:"收货人姓名" json:"contactName"`          // 收货人姓名
	ContactMobile string `validate:"required,mobile" label:"收货人电话" json:"contactMobile"` // 收货人电话
}
