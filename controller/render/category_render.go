// Package render
// @Author:        asus
// @Description:   $
// @File:          category_render
// @Data:          2022/2/2816:48
//
package render

import (
	"gorm.io/gorm"
	"new-project/models"
)

type Category struct {
	ID           uint           `json:"id" example:"1"`
	CreatedAt    uint32         `json:"createdAt" example:"1646036184"`
	UpdatedAt    uint32         `json:"updatedAt" example:"1646036184"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt" swaggertype:"string" example:"2022-02-01 16:51:21"`
	Path         string         `json:"path" example:"-"`
	CategoryName string         `json:"categoryName" example:"一级分类"`
	CategoryImg  string         `json:"categoryImg" example:"http://test/image/1.jpg"`
	IsFinal      bool           `json:"isFinal" example:"false"`
	Sort         uint           `json:"sort" example:"0"`
	CategoryID   uint           `json:"categoryID" example:"0"`
	Level        uint           `json:"level" example:"1"`
}

func BuildCreategory(category *models.Category) *Category {
	if category == nil {
		return nil
	}

	return &Category{
		ID:           category.ID,
		CreatedAt:    category.CreatedAt,
		UpdatedAt:    category.UpdatedAt,
		DeletedAt:    category.DeletedAt,
		Path:         category.Path,
		CategoryName: category.CategoryName,
		CategoryImg:  category.CategoryImg,
		IsFinal:      category.IsFinal,
		Sort:         category.Sort,
		CategoryID:   category.CategoryID,
		Level:        category.Level,
	}
}

func BuildCreategories(categories []*models.Category) []*Category {
	list := make([]*Category, 0)
	if len(categories) < 1 {
		return list
	}

	for _, category := range categories {
		list = append(list, BuildCreategory(category))
	}
	return list
}
