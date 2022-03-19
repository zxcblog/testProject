package render

import (
	"encoding/json"
	"new-project/models"
)

type ProductSku struct {
	ID           uint                   `json:"id"`           //ID
	ProductId    uint                   `json:"productId"`    //商品ID
	SkuAttribute map[string]interface{} `json:"skuAttribute"` //商品规格组合数据
	Stock        uint                   `json:"stock"`        //库存
	Price        uint                   `json:"price"`        //价格
}

func BuildProductSku(productSku *models.ProductSku) *ProductSku {
	if productSku == nil {
		return nil
	}
	var skuAttributeSlide map[string]interface{}

	json.Unmarshal([]byte(productSku.SkuAttribute), &skuAttributeSlide)

	return &ProductSku{
		ID:           productSku.ID,
		ProductId:    productSku.ProductId,
		SkuAttribute: skuAttributeSlide,
		Stock:        productSku.Stock,
		Price:        productSku.Price,
	}
}

func BuildProductSkuList(productSkuList []*models.ProductSku) []*ProductSku {
	list := make([]*ProductSku, 0)
	if len(productSkuList) < 1 {
		return list
	}

	for _, productSku := range productSkuList {
		list = append(list, BuildProductSku(productSku))
	}
	return list
}
