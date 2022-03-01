package cache

import (
	"new-project/global"
	"new-project/models"
	"new-project/repositories"
)

var CategoryCache = map[uint]*models.Category{}

func GetCategoryByID(id uint) *models.Category {
	if _, ok := CategoryCache[id]; !ok {
		CategoryCache[id] = repositories.CategoryRepositories.Get(global.DB, id)
	}

	return CategoryCache[id]
}

func AddCategory(category *models.Category) {
	if _, ok := CategoryCache[category.ID]; !ok {
		CategoryCache[category.ID] = category
	}
}

func UpdateCategory(category *models.Category) {
	CategoryCache[category.ID] = category
}

func DelCategory(id uint) {
	delete(CategoryCache, id)
}
