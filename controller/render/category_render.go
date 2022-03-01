// Package render
// @Author:        asus
// @Description:   $
// @File:          category_render
// @Data:          2022/2/2816:48
//
package render

import (
	"new-project/cache"
	"new-project/models"
	"strconv"
	"strings"
)

type Category struct {
	ID           uint   `json:"id" example:"1"`
	PathName     string `json:"pathName,omitempty" example:""`
	CategoryName string `json:"categoryName" example:"一级分类"`
	CategoryImg  string `json:"categoryImg" example:"http://test/image/1.jpg"`
	IsFinal      bool   `json:"isFinal" example:"false"`
	Sort         uint   `json:"sort" example:"0"`
	CategoryID   uint   `json:"categoryID" example:"0"`
}

func BuildCreategory(category *models.Category) *Category {
	if category == nil {
		return nil
	}

	return &Category{
		ID:           category.ID,
		CategoryName: category.CategoryName,
		CategoryImg:  category.CategoryImg,
		IsFinal:      category.IsFinal,
		Sort:         category.Sort,
		CategoryID:   category.CategoryID,
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

func BuildCreategoryPathName(categories []*models.Category) []*Category {
	list := make([]*Category, 0)

	for _, category := range categories {
		ids := strings.Split(category.Path, "-")
		if len(ids) < 3 {
			continue
		}

		path := ""
		if id, err := strconv.Atoi(ids[1]); err != nil {
			path += cache.GetCategoryByID(uint(id)).CategoryName + "/"
		}

		if id, err := strconv.Atoi(ids[2]); err != nil {
			path += cache.GetCategoryByID(uint(id)).CategoryName + "/"
		}

		item := BuildCreategory(category)
		item.PathName = path + category.CategoryName
		list = append(list, item)
	}
	return list
}
