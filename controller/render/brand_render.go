// Package render
// @Author:        asus
// @Description:   $
// @File:          category_render
// @Data:          2022/2/2816:48
//
package render

import (
	"new-project/models"
)

type Brand struct {
	ID          uint   `json:"id" example:"1"`
	CategoryID  uint   `json:"categoryID" example:"1"`
	BrandName   string `json:"brandName" example:"华硕"`
	BrandImg    string `json:"brandImg" example:""`
	Description string `json:"description" example:"华硕描述"`
}

func BuildBrand(brand *models.Brand) *Brand {
	if brand == nil {
		return nil
	}

	return &Brand{
		ID:          brand.ID,
		CategoryID:  brand.CategoryID,
		BrandName:   brand.BrandName,
		BrandImg:    brand.BrandImg,
		Description: brand.Description,
	}
}

func BuildBrands(brands []*models.Brand) []*Brand {
	list := make([]*Brand, 0, len(brands))
	if len(brands) < 1 {
		return list
	}

	for _, brand := range brands {
		list = append(list, BuildBrand(brand))
	}
	return list
}
