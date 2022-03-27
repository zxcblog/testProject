// Package models
// @Author:        asus
// @Description:   $
// @File:          model
// @Data:          2022/2/2611:53
//
package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"new-project/pkg/config"
)

type Model struct {
	ID        uint           `gorm:"primary_key"`
	CreatedAt uint32         `gorm:"autoCreateTime"`
	UpdatedAt uint32         `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NewDBEngine(option *config.Database, models ...interface{}) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=Local",
		option.Username, option.Password, option.Host, option.Port, option.DBName, option.Charset, option.ParseTime,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   option.TablePrefix, // 表前缀
			SingularTable: true,               // 使用单表复数名
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(option.MaxIdleConns) // SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(option.MaxOpenConns) // SetMaxOpenConns 设置打开数据库连接的最大数量。

	if err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models...); err != nil {
		return nil, fmt.Errorf("自动化生成表失败：%s", err.Error())
	}

	return db, nil
}
