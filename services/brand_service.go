// Package services
// @Author:        asus
// @Description:   $
// @File:          category_service
// @Data:          2022/2/2815:51
//
package services

import (
	"go.uber.org/zap"
	"new-project/global"
	"new-project/models"
	"new-project/pkg/errcode"
	"new-project/repositories"
)

var BrandService = NewBrandService()

type brandService struct{}

func NewBrandService() *brandService {
	return &brandService{}
}

// Get 获取品牌信息
func (b *brandService) Get(id uint) *models.Brand {
	return repositories.BrandRepositories.Get(global.DB, id)
}

// GetListPage 获取品牌列表
func (b *brandService) GetListPage(catrgoryID uint, page, pageSize int) ([]*models.Brand, int64) {
	db := global.DB
	if catrgoryID > 0 {
		db.Where("category_id", catrgoryID)
	}

	return repositories.BrandRepositories.GetList(db, page, pageSize)
}

// Create 创建品牌
func (b *brandService) Create(brand *models.Brand) error {
	if err := b.setCategory(brand.CategoryID); err != nil {
		return err
	}

	err := repositories.BrandRepositories.Create(global.DB, brand)
	if err != nil {
		global.Logger.Error("品牌创建失败", zap.Error(err))
		return errcode.CreateError.SetMsg("品牌创建失败")
	}
	return nil
}

// Update 修改分类
func (b *brandService) Update(brand *models.Brand) error {
	if err := b.setCategory(brand.CategoryID); err != nil {
		return err
	}

	err := repositories.BrandRepositories.Update(global.DB, brand)
	if err != nil {
		global.Logger.Error("品牌修改失败", zap.Error(err))
		return errcode.CreateError.SetMsg("品牌修改失败")
	}
	return nil
}

func (b *brandService) setCategory(categoryID uint) error {
	category := CategoryService.Get(categoryID)
	if category == nil {
		return errcode.NotFound
	}
	if !category.IsFinal {
		return errcode.RequestError.SetMsg("所选分类不是最终类")
	}
	return nil
}

// Delete 删除分类
func (b *brandService) Delete(id uint) error {
	err := repositories.BrandRepositories.Delete(global.DB, id)
	if err != nil {
		global.Logger.Error("品牌删除失败", zap.Error(err))
		return errcode.DelError
	}
	return nil
}
