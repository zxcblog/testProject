package repositories

import (
	"gorm.io/gorm"
	"new-project/models"
)

var BrandRepositories = NewBrandRepositories()

type brandRepositories struct{}

func NewBrandRepositories() *brandRepositories {
	return &brandRepositories{}
}

// Get 通过id进行搜索
func (b *brandRepositories) Get(db *gorm.DB, id uint) *models.Brand {
	ret := &models.Brand{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}

// GetList 获取分类列表
func (b *brandRepositories) GetList(db *gorm.DB, page, pageSize int) ([]*models.Brand, int64) {
	list := make([]*models.Brand, 0, pageSize)
	var total int64
	db.Model(models.Brand{}).Count(&total).
		Order("sort desc").
		Limit(pageSize).Offset((page - 1) * pageSize).Find(&list)
	return list, total
}

// Create 创建
func (b *brandRepositories) Create(db *gorm.DB, brand *models.Brand) error {
	return db.Create(brand).Error
}

// Update 修改
func (b *brandRepositories) Update(db *gorm.DB, brand *models.Brand) error {
	return db.Save(brand).Error
}

// Delete 通过id删除
func (b *brandRepositories) Delete(db *gorm.DB, id uint) error {
	return db.Delete(&models.Brand{}, "id = ?", id).Error
}
