// Package global
// @Author:        asus
// @Description:   $
// @File:          global
// @Data:          2022/2/2613:02
//
package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"io"
	"new-project/pkg/app"
)

var AccessLog io.Writer
var Logger *zap.Logger
var DB *gorm.DB
var Validate *app.Translations
