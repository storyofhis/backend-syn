package service

import (
	"context"

	"github.com/storyofhis/category-svc/internal/repository/gorm"
	"github.com/storyofhis/category-svc/internal/repository/model"
)

type categoryService struct {
	repo *gorm.CategoryRepo
}

type CategoryService interface {
	CreateCategory(ctx context.Context, category *model.Category) (*model.Category, error)
	GetCategoryById(ctx context.Context, id string) (*model.Category, error)
	GetCategories(ctx context.Context) ([]*model.Category, error)
}

func NewCategoryService(repo *gorm.CategoryRepo) CategoryService {
	return &categoryService{
		repo: repo,
	}
}

func (s *categoryService) CreateCategory(ctx context.Context, category *model.Category) (*model.Category, error) {
	return s.repo.CreateCategory(ctx, category)
}

func (s *categoryService) GetCategoryById(ctx context.Context, id string) (*model.Category, error) {
	return s.repo.GetCategoryById(ctx, id)
}

func (s *categoryService) GetCategories(ctx context.Context) ([]*model.Category, error) {
	categories, err := s.repo.GetCategories(ctx, nil)
	if err != nil {
		return nil, err
	}

	var categoriesPointer []*model.Category
	for _, category := range categories {
		categoriesPointer = append(categoriesPointer, category)
	}
	return categoriesPointer, nil
}
