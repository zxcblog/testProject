// Package models
// @Author:        asus
// @Description:   $
// @File:          shop
// @Data:          2022/2/2613:07
//
package models

var Models = []interface{}{
	&Product{}, &User{}, &Category{}, &ProductSku{}, &Brand{}, &Upload{}, &ProductSkuKey{}, &ProductSkuValue{},
}

// User 用户表
type User struct {
	Model
	Status   uint   `gorm:"type:tinyint(1);index:idx_shop_user_status;default:1;not null;comment:状态1正常2非正常"`
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

// Brand 品牌表
type Brand struct {
	Model
	BrandName   string `gorm:"size:512;comment:品牌名称"`
	BrandImg    string `gorm:"type:text;comment:品牌图片"`
	Description string `gorm:"type:text;comment:描述"`
	CategoryID  uint   `gorm:"index:idx_category_shop_category_id;default:0;comment:所属分类"`
}

// Product 商品表
type Product struct {
	Model
	ProductName  string `gorm:"size:50;comment:商品名称"`
	ProductTitle string `gorm:"size:100;comment:商品标题"`
	CategoryID   uint   `gorm:"index:idx_category_shop_category_id;default:0;comment:分类ID"`
	BrandId      uint   `gorm:"index:idx_category_shop_brand_id;default:0;comment:品牌ID"`
	CategoryImg  string `gorm:"size:100;comment:商品主图"`
	PicImg       string `gorm:"type:text;comment:商品轮播图"`
	Status       uint   `gorm:"type:tinyint(1);default:1;comment:上架状态:1=下架,2=申请上架,3=上架"`
	Introduction string `gorm:"type:text;comment:详情介绍"`
}

// ProductSku 商品sku表
type ProductSku struct {
	Model
	ProductId    uint   `gorm:"index:idx_sku_shop_product_id;comment:商品id"`
	SkuAttribute string `gorm:"type:text;comment:商品规格组合数据"`
	Stock        uint   `gorm:"dedault:1;comment:库存"`
	Price        uint   `gorm:"dedault:1;comment:价格"`
}

// ProductSku 商品sku属性key表
type ProductSkuKey struct {
	Model
	ProductId    uint   `gorm:"index:idx_sku_key_shop_product_id;comment:商品id"`
	BrandId      uint   `gorm:"index:idx_sku_key_shop_brand_id;default:0;comment:品牌ID"`
	AttributeKey string `gorm:"size:100;comment:属性值"`
}

// ProductSku 商品sku属性value表
type ProductSkuValue struct {
	Model
	ProductId       uint   `gorm:"index:idx_sku_value_shop_product_id;comment:商品id"`
	ProductSkuKeyId uint   `gorm:"index:idx_sku_value_shop_product_id;comment:属性id"`
	AttributeValue  string `gorm:"size:100;comment:属性值"`
	Sort            uint   `gorm:"comment:排序"`
}

// Upload 文件上传保存
type Upload struct {
	Model
	FileSize uint   `gorm:"default:0;comment:文件大小"`
	SavePath string `gorm:"size:512;comment:文件保存地址"`
	OldName  string `gorm:"size:512;comment:文件原名称"`
	NewName  string `gorm:"size:512;comment:新名称"`
	FileType string `gorm:"size:512;comment:文件类型"`
}
