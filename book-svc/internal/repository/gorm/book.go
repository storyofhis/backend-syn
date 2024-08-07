package gorm

import (
	"context"

	"github.com/storyofhis/book-svc/internal/repository/model"
	"gorm.io/gorm"
)

type BookRepo struct {
	db *gorm.DB
}

func NewBookRepo(db *gorm.DB) *BookRepo {
	return &BookRepo{
		db: db,
	}
}

func (repo *BookRepo) CreateBook(ctx context.Context, book *model.Book) (*model.Book, error) {
	if err := repo.db.WithContext(ctx).Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (repo *BookRepo) GetBookById(ctx context.Context, id string) (*model.Book, error) {
	book := new(model.Book)
	if err := repo.db.WithContext(ctx).Where("id = ?", id).Take(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

// GetBooks fetches multiple books based on given criteria
func (repo *BookRepo) GetBooks(ctx context.Context, criteria map[string]interface{}) ([]model.Book, error) {
	var books []model.Book
	query := repo.db.WithContext(ctx).Model(&model.Book{})

	for key, value := range criteria {
		query = query.Where(key+" = ?", value)
	}

	if err := query.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}
