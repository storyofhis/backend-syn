package gorm

import (
	"context"

	"github.com/storyofhis/category-svc/internal/repository/model"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{db: db}
}

func (repo *CategoryRepo) CreateCategory(ctx context.Context, category *model.Category) (*model.Category, error) {
	if err := repo.db.WithContext(ctx).Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (repo *CategoryRepo) GetCategoryById(ctx context.Context, id string) (*model.Category, error) {
	category := new(model.Category)
	if err := repo.db.WithContext(ctx).Where("id = ?", id).Take(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (repo *CategoryRepo) GetCategories(ctx context.Context, criteria map[string]interface{}) ([]*model.Category, error) {
	var categories []*model.Category
	query := repo.db.WithContext(ctx).Model(&model.Category{})

	for key, value := range criteria {
		query = query.Where(key+" = ?", value)
	}

	if err := query.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
