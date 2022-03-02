// Package global
// @Author:        asus
// @Description:   $
// @File:          global
// @Data:          2022/2/2613:02
//
package global

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"new-project/pkg/app"
)

var (
	AccessLog *zap.Logger
	Logger    *zap.Logger
	DB        *gorm.DB
	Validate  *app.Translations
	Redis     *redis.Client
)
