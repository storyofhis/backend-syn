package gorm

import (
	"context"

	"github.com/storyofhis/author-service/internal/repository/model"
	"gorm.io/gorm"
)

type AuthorRepo struct {
	db *gorm.DB
}

func NewAuthorRepo(db *gorm.DB) *AuthorRepo {
	return &AuthorRepo{
		db: db,
	}
}

func (repo *AuthorRepo) CreateAuthor(ctx context.Context, author *model.Author) (*model.Author, error) {
	if err := repo.db.WithContext(ctx).Create(author).Error; err != nil {
		return nil, err
	}
	return author, nil
}

func (repo *AuthorRepo) GetAuthorById(ctx context.Context, id string) (*model.Author, error) {
	author := new(model.Author)
	if err := repo.db.WithContext(ctx).Where("id = ?", id).Take(author).Error; err != nil {
		return nil, err
	}
	return author, nil
}

func (repo *AuthorRepo) GetAuthors(ctx context.Context, criteria map[string]interface{}) ([]*model.Author, error) {
	var authors []*model.Author
	query := repo.db.WithContext(ctx).Model(&model.Author{})

	for key, value := range criteria {
		query = query.Where(key+" = ?", value)
	}
	if err := query.Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}
