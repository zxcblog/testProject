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

// Get 获取分类信息
func (c *categoryService) Get(id uint) *models.Category {
	return repositories.CategoryRepositories.Get(global.DB, id)
}

// GetListPage 获取分类列表
func (c *categoryService) GetListPage(db *gorm.DB, page, pageSize int) ([]*models.Category, int64) {
	return repositories.CategoryRepositories.GetList(db, page, pageSize)
}

// Create 创建分类
func (c *categoryService) Create(category *models.Category) error {
	if err := c.setCategoryParent(category); err != nil {
		return err
	}

	err := repositories.CategoryRepositories.Create(global.DB, category)
	if err != nil {
		global.Logger.Error("分类创建失败", zap.Error(err))
		return errcode.CreateError.SetMsg("分类创建失败")
	}

	return nil
}

// Update 修改分类
func (c *categoryService) Update(category *models.Category) error {
	if err := c.setCategoryParent(category); err != nil {
		return err
	}

	err := repositories.CategoryRepositories.Update(global.DB, category)
	if err != nil {
		global.Logger.Error("分类修改失败", zap.Error(err))
		return errcode.CreateError.SetMsg("分类修改失败")
	}
	return nil
}

func (c *categoryService) setCategoryParent(category *models.Category) error {
	if category.CategoryID == 0 {
		return nil
	}

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
	return nil
}

// Delete 删除分类
func (c *categoryService) Delete(id uint) error {
	category := repositories.CategoryRepositories.Get(global.DB, id)
	if category == nil {
		return errcode.SelectError
	}

	err := global.DB.Transaction(func(tx *gorm.DB) error {
		err := repositories.CategoryRepositories.Delete(tx, id)
		if err != nil {
			return err
		}

		return repositories.CategoryRepositories.DeleteWhere(tx.Where("path like ?", category.Path+"%"))
	})
	if err != nil {
		global.Logger.Error("分类删除失败", zap.Error(err))
		return errcode.TransactionError.SetMsg("分类删除失败")
	}

	return nil
}
