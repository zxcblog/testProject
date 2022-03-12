package repositories

import (
	"gorm.io/gorm"
	"new-project/models"
)

var UploadRepositories = NewUploadRepositories()

type uploadRepositories struct{}

func NewUploadRepositories() *uploadRepositories {
	return &uploadRepositories{}
}

//// Get 通过id进行搜索
//func (b *uploadRepositories) Get(db *gorm.DB, id uint) *models.Upload {
//	ret := &models.Upload{}
//	if err := db.First(ret, "id = ?", id).Error; err != nil {
//		return nil
//	}
//
//	return ret
//}
//
//// GetList 获取分类列表
//func (b *uploadRepositories) GetList(db *gorm.DB, page, pageSize int) ([]*models.Upload, int64) {
//	list := make([]*models.Upload, pageSize)
//	var total int64
//	db.Model(models.Upload{}).Count(&total).
//		Order("sort desc").
//		Limit(pageSize).Offset((page - 1) * pageSize).Find(&list)
//	return list, total
//}

// Create 创建
func (*uploadRepositories) Create(db *gorm.DB, upload *models.Upload) error {
	return db.Create(upload).Error
}

//// Update 修改
//func (b *uploadRepositories) Update(db *gorm.DB, upload *models.Upload) error {
//	return db.Save(upload).Error
//}
//
//// Delete 通过id删除
//func (b *uploadRepositories) Delete(db *gorm.DB, id uint) error {
//	return db.Delete(&models.Upload{}, "id = ?", id).Error
//}
