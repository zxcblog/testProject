package render

import (
	"encoding/json"
	"new-project/models"
)

type Product struct {
	ID            uint   `json:"id"`            //ID
	ProductName   string `json:"productName"`   //商品名称
	ProductTitle  string `json:"productTitle"`  //商品标题
	CategoryID    uint   `json:"categoryID"`    //商品分类
	BrandId       uint   `json:"brandId"`       //品牌分类
	CategoryImgId uint   `json:"categoryImgId"` //商品主图ID
	PicImgIds     []uint `json:"picImgIds"`     //商品缩略图
	Introduction  string `json:"introduction"`  //商品详情
}

func BuildProduct(product *models.Product) *Product {
	if product == nil {
		return nil
	}
	var picImgIdsSlide []uint

	json.Unmarshal([]byte(product.PicImgIds), &picImgIdsSlide)

	return &Product{
		ID:            product.ID,
		ProductName:   product.ProductName,
		ProductTitle:  product.ProductTitle,
		CategoryID:    product.CategoryID,
		BrandId:       product.BrandId,
		CategoryImgId: product.CategoryImgId,
		PicImgIds:     picImgIdsSlide,
		Introduction:  product.Introduction,
	}
}
