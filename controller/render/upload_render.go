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

type Upload struct {
	ID       uint   `json:"id" example:"1"`
	FileSize uint   `gorm:"default:0;comment:文件大小"`
	SavePath string `gorm:"size:512;comment:文件保存地址"`
	OldName  string `gorm:"size:512;comment:文件原名称"`
	NewName  string `gorm:"size:512;comment:新名称"`
	FileType string `gorm:"size:512;comment:文件类型"`
	FileExt  string `gorm:"size:512;comment:文件后缀"`
}

func BuildUpload(upload *models.Upload) *Upload {
	if upload == nil {
		return nil
	}

	return &Upload{
		ID:       upload.ID,
		FileSize: upload.FileSize,
		SavePath: upload.SavePath,
		OldName:  upload.OldName,
		NewName:  upload.NewName,
		FileType: upload.FileType,
		FileExt:  upload.FileExt,
	}
}

//
//func BuildBrands(brands []*models.Brand) []*Brand {
//	list := make([]*Brand, 0, len(brands))
//	if len(brands) < 1 {
//		return list
//	}
//
//	for _, brand := range brands {
//		list = end(list, BuildBrand(brand))
//	}
//	return list
//}
