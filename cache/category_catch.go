package cache

import (
	"context"
	"new-project/global"
	"new-project/models"
	"new-project/pkg/util"
	"new-project/repositories"
	"strconv"
)

var CategoryCache = NewCategoryCache(context.Background(), "category")

type categoryCatch struct {
	ctx context.Context
	Key string
}

func NewCategoryCache(ctx context.Context, key string) *categoryCatch {
	return &categoryCatch{ctx: ctx, Key: key}
}

func (c *categoryCatch) SetContext(ctx context.Context) *categoryCatch {
	newC := *c
	newC.ctx = ctx
	return &newC
}

// GetCategoryByID 通过id获取到redis缓存
func (c *categoryCatch) GetCategoryByID(id uint) *models.Category {
	field := strconv.Itoa(int(id))
	str, err := global.Redis.HGet(c.ctx, c.Key, field).Result()
	if err != nil {
		return nil
	}

	category := &models.Category{}
	if str != "" {
		util.StringToStruct(str, &category)
		if category != nil {
			return category
		}
	}

	category = repositories.CategoryRepositories.Get(global.DB, id)
	if category == nil {
		return nil
	}

	c.AddCategory(category)
	return category
}

func (c *categoryCatch) AddCategory(category *models.Category) (bool, error) {
	field := strconv.Itoa(int(category.ID))

	flag, err := global.Redis.HExists(c.ctx, c.Key, field).Result()
	if err != nil {
		return false, err
	}

	// 如果为true, 说明已存在, 不进行添加
	if flag {
		return false, nil
	}

	// 添加缓存
	res, err := global.Redis.HSet(c.ctx, c.Key, field, util.StructToString(category)).Result()
	if err != nil {
		return false, err
	}

	return res > 0, nil
}

func (c *categoryCatch) UpdateCategory(category *models.Category) (bool, error) {
	field := strconv.Itoa(int(category.ID))

	flag, err := global.Redis.HExists(c.ctx, c.Key, field).Result()
	if err != nil {
		return false, err
	}

	// 如果不存在， 不进行处理
	if !flag {
		return false, nil
	}

	res, err := global.Redis.HSet(c.ctx, c.Key, field, util.StructToString(category)).Result()
	if err != nil {
		return false, err
	}

	return res > 0, nil
}

// DelCategory 删除缓存
func (c *categoryCatch) DelCategory(ids ...uint) {
	for _, id := range ids {
		global.Redis.HDel(c.ctx, c.Key, strconv.Itoa(int(id)))
	}
}
