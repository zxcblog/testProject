// Package services
// @Author:        asus
// @Description:   $
// @File:          category_service
// @Data:          2022/2/2815:51
//
package services

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"new-project/global"
	"new-project/models"
	"new-project/pkg/errcode"
	"new-project/repositories"
	"strconv"
)

var CategoryService = NewCategoryService()

type categoryService struct{}

func NewCategoryService() *categoryService {
	return &categoryService{}
}

// Create 创建分类
func (c *categoryService) Create(category *models.Category) error {
	if category.CategoryID != 0 {
		parentCategory := repositories.CategoryRepositories.Get(global.DB, category.CategoryID)
		if parentCategory == nil {
			return errcode.SelectError.SetMsg("所属分类不存在")
		}
		if parentCategory.IsFinal {
			return errcode.RequestError.SetMsg("当前分类为最终类，不能添加子类目")
		}
		if parentCategory.Level >= 2 && !category.IsFinal {
			return errcode.RequestError.SetMsg("当前分类必须为最终类")
		}
		category.Path = parentCategory.Path + strconv.Itoa(int(parentCategory.ID)) + "-"
		category.Level = parentCategory.Level + 1
	}

	err := repositories.CategoryRepositories.Create(global.DB, category)
	if err != nil {
		global.Logger.Error("分类创建失败", zap.Error(err))
		return errcode.CreateError.SetMsg("分类创建失败")
	}

	return nil
}

func (c *categoryService) GetListPage(db *gorm.DB, page, pageSize int) ([]*models.Category, int64) {
	return repositories.CategoryRepositories.GetList(db, page, pageSize)
}
