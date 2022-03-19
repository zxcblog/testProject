package render

import (
	"encoding/json"
	"new-project/models"
)

type ProductSkuKeyValue struct {
	ID             uint     `json:"id"`             //ID
	ProductId      uint     `json:"productId"`      //商品ID
	BrandId        uint     `json:"brandId"`        //分类ID
	AttributeKey   string   `json:"attributeKey"`   //属性key值
	AttributeValue []string `json:"attributeValue"` //属性value
}

func BuildProductSkuKeyValue(productSkuKeyValue *models.ProductSkuKeyValue) *ProductSkuKeyValue {
	if productSkuKeyValue == nil {
		return nil
	}
	var skuAttributeValueSlide []string

	json.Unmarshal([]byte(productSkuKeyValue.AttributeValue), &skuAttributeValueSlide)

	return &ProductSkuKeyValue{
		ID:             productSkuKeyValue.ID,
		ProductId:      productSkuKeyValue.ProductId,
		BrandId:        productSkuKeyValue.BrandId,
		AttributeKey:   productSkuKeyValue.AttributeKey,
		AttributeValue: skuAttributeValueSlide,
	}
}

func BuildProductSkuKeyValueList(productSkuKeyValueList []*models.ProductSkuKeyValue) []*ProductSkuKeyValue {
	list := make([]*ProductSkuKeyValue, 0)
	if len(productSkuKeyValueList) < 1 {
		return list
	}

	for _, productSkuKeyValue := range productSkuKeyValueList {
		list = append(list, BuildProductSkuKeyValue(productSkuKeyValue))
	}
	return list
}
