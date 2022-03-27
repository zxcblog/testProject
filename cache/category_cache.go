package cache

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"new-project/global"
	"new-project/models"
	"new-project/pkg/util"
	"new-project/repositories"
	"strconv"
	"strings"
)

var CategoryCache = NewCategoryCache(context.Background(), "category:")

type categoryCache struct {
	ctx context.Context
	Key string
}

func NewCategoryCache(ctx context.Context, key string) *categoryCache {
	return &categoryCache{ctx: ctx, Key: key}
}

func (c *categoryCache) SetContext(ctx context.Context) *categoryCache {
	newC := *c
	newC.ctx = ctx
	return &newC
}

// GetCategoryByID 通过id获取到redis缓存
func (c *categoryCache) GetCategoryByID(id uint) *models.Category {
	field := c.Key + "*" + strconv.Itoa(int(id))
	keys, err := global.Redis.Keys(c.ctx, field).Result()
	if err != nil {
		global.Logger.Error("分类获取失败：", zap.Error(err))
	}

	category := &models.Category{}
	if len(keys) > 0 {
		res, _ := global.Redis.Get(c.ctx, keys[0]).Result()
		util.StringToStruct(res, &category)
		if category.ID > 0 {
			return category
		}
	}

	// 返回数据
	category = repositories.CategoryRepositories.Get(global.DB, id)
	if category != nil {
		c.SetCategory(category)
	}
	return category
}

//GetCategoryByPath 通过路径获取分类信息
func (c *categoryCache) GetCategoryByPath(path string) *models.Category {
	if path == "" {
		return nil
	}

	key := c.Key + path
	category := &models.Category{}
	res, _ := global.Redis.Get(c.ctx, key).Result()
	util.StringToStruct(res, &category)
	if category.ID > 0 {
		return category
	}

	// 返回数据
	ids := strings.Split(strings.Trim(path, models.CategorySep), models.CategorySep)

	id, _ := strconv.Atoi(ids[len(ids)-1])
	category = repositories.CategoryRepositories.Get(global.DB, uint(id))
	if category != nil {
		c.SetCategory(category)
	}
	return category
}

// SetCategory 设置分类缓存
func (c *categoryCache) SetCategory(category *models.Category) (bool, error) {
	field := c.Key + category.Path + strconv.Itoa(int(category.ID))

	// 添加缓存
	res, err := global.Redis.Set(c.ctx, field, util.StructToString(category), 0).Result()
	if err != nil {
		global.Logger.Error("分类缓存失败：", zap.Error(err))
		return false, err
	}

	return res != "", nil
}

// DelCategory 删除缓存
func (c *categoryCache) DelCategory(key string) (int64, error) {
	field := c.Key + key + "*"
	fmt.Println(field)
	keys, err := global.Redis.Keys(c.ctx, field).Result()
	if err != nil {
		global.Logger.Error("分类删除失败：", zap.Error(err))
		return 0, err
	}

	return global.Redis.Del(c.ctx, keys...).Result()
}
