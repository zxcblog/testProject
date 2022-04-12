// Package services
// @Author:        asus
// @Description:   $
// @File:          category_service
// @Data:          2022/2/2815:51
//
package services

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"new-project/cache"
	"new-project/global"
	"new-project/models"
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
	return cache.CategoryCache.GetCategoryByID(id)
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
		return errors.New("分类创建失败")
	}

	_, err = cache.CategoryCache.SetCategory(category)
	if err != nil {
		global.Logger.Error("分类缓存处理错误", zap.Error(err))
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
		return errors.New("分类修改失败")
	}

	_, err = cache.CategoryCache.SetCategory(category)
	if err != nil {
		global.Logger.Error("分类缓存处理错误", zap.Error(err))
	}
	return nil
}

//setCategoryParent 设置分类的path前缀及其他属性
func (c *categoryService) setCategoryParent(category *models.Category) error {
	if category.CategoryID == 0 {
		return nil
	}

	parentCategory := repositories.CategoryRepositories.Get(global.DB, category.CategoryID)
	if parentCategory == nil {
		return errors.New("所属分类不存在")
	}
	if parentCategory.IsFinal {
		return errors.New("当前分类为最终类，不能添加子类目")
	}
	if parentCategory.Level >= 2 {
		category.IsFinal = true
	}
	category.Path = parentCategory.Path + strconv.Itoa(int(parentCategory.ID)) + models.CategorySep
	category.Level = parentCategory.Level + 1
	return nil
}

// Delete 删除分类
func (c *categoryService) Delete(id uint) error {
	category := c.Get(id)
	if category == nil {
		return nil
	}

	// 要被删除的子集
	path := category.Path + strconv.Itoa(int(category.ID))
	err := global.DB.Transaction(func(tx *gorm.DB) error {
		err := repositories.CategoryRepositories.Delete(tx, id)
		if err != nil {
			return err
		}

		// 删除分类时， 同时删除子分类

		return repositories.CategoryRepositories.DeleteWhere(tx.Where("path like ?", path+"%"))
	})
	if err != nil {
		global.Logger.Error("分类删除失败", zap.Error(err))
		return errors.New("分类删除失败")
	}

	//删除分类时， 删除缓存信息
	cache.CategoryCache.DelCategory(path)
	return nil
}

// QueryName 通过名称搜索分类
func (c *categoryService) QueryName(categoryName string) (list []*models.Category) {
	db := global.DB.Where("category_name like ?", "%"+categoryName+"%").Where("is_final", true)

	var err error
	list, err = repositories.CategoryRepositories.GetWhereList(db)
	if err != nil {
		global.Logger.Error("分类名称模糊查询失败", zap.Error(err))
	}

	return list
}
