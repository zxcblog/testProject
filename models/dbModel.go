// Package models
// @Author:        asus
// @Description:   $
// @File:          shop
// @Data:          2022/2/2613:07
//
package models

var Models = []interface{}{
	&Product{}, &User{}, &Category{}, &ProductSku{}, &Brand{}, &Upload{}, &ProductSkuKeyValue{}, &Address{},
}

// User 用户表
type User struct {
	Model
	Status      uint   `gorm:"type:tinyint(1);index:idx_shop_user_status;default:1;not null;comment:状态1正常2非正常"`
	Username    string `gorm:"size:32;unique;comment:账号"`
	Nickname    string `gorm:"size:16;comment:昵称"`
	Avatar      string `gorm:"type:text;comment:头像"`
	Password    string `gorm:"size:512;comment:密码"`
	Salt        string `gorm:"size:16;comment:密码加盐"`
	AccountType uint   `gorm:"default:1;comment:账号类型1前台用户2后台管理员"`
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
	ProductName   string `gorm:"size:50;comment:商品名称"`
	ProductTitle  string `gorm:"size:100;comment:商品标题"`
	CategoryID    uint   `gorm:"index:idx_category_shop_category_id;default:0;comment:分类ID"`
	BrandId       uint   `gorm:"index:idx_category_shop_brand_id;default:0;comment:品牌ID"`
	CategoryImgId uint   `gorm:"not null;comment:商品主图ID"`
	PicImgIds     string `gorm:"type:text;comment:商品轮播图ID组"`
	Status        uint   `gorm:"type:tinyint(1);default:1;comment:上架状态:1=下架,2=申请上架,3=上架"`
	Introduction  string `gorm:"type:text;comment:详情介绍"`
}

// ProductSku 商品sku表
type ProductSku struct {
	Model
	ProductId    uint   `gorm:"index:idx_sku_shop_product_id;comment:商品id"`
	SkuAttribute string `gorm:"type:text;comment:商品规格组合数据"`
	Stock        uint   `gorm:"default:1;comment:库存"`
	Price        uint   `gorm:"default:1;comment:价格"`
}

// ProductSkuKeyValue 商品sku属性key和value表
type ProductSkuKeyValue struct {
	Model
	ProductId      uint   `gorm:"index:idx_sku_key_shop_product_id;comment:商品id"`
	BrandId        uint   `gorm:"index:idx_sku_key_shop_brand_id;default:0;comment:品牌ID"`
	AttributeKey   string `gorm:"size:100;comment:属性key值"`
	AttributeValue string `gorm:"size:100;comment:属性value json值"`
}

// Upload 文件上传保存
type Upload struct {
	Model
	UserId   uint   `gorm:"index:idx_upload_user_id;default:0;comment:用户id"`
	FileSize uint   `gorm:"default:0;comment:文件大小"`
	SavePath string `gorm:"size:512;comment:文件保存地址"`
	OldName  string `gorm:"size:512;comment:文件原名称"`
	NewName  string `gorm:"size:512;comment:新名称"`
	FileType string `gorm:"size:512;comment:文件类型"`
	FileExt  string `gorm:"size:512;comment:文件后缀"`
}

// Address 用户地址管理
type Address struct {
	Model
	IsDefault     bool   `gorm:"type:tinyint(1);default:0;comment:默认地址:1=默认地址"`
	UserId        uint   `gorm:"index:idx_address_user_id;default:0;comment:用户id"`
	Province      string `gorm:"size:512;comment:省"`
	City          string `gorm:"size:512;comment:市"`
	Area          string `gorm:"size:512;comment:区"`
	Street        string `gorm:"size:512;comment:街道"`
	Desc          string `gorm:"size:512;comment:详细地址"`
	ContactName   string `gorm:"size:512;comment:收货人姓名"`
	ContactMobile string `gorm:"size:512;comment:收货人电话"`
}
