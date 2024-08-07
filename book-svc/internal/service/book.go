package service

import (
	"context"

	"github.com/storyofhis/book-svc/internal/repository/gorm"
	"github.com/storyofhis/book-svc/internal/repository/model"
)

type bookService struct {
	repo gorm.BookRepo
}

type BookService interface {
	CreateBook(ctx context.Context, book *model.Book) (*model.Book, error)
	GetBookById(ctx context.Context, id string) (*model.Book, error)
	ListBooks(ctx context.Context) ([]*model.Book, error)
}

func NewBookService(repo gorm.BookRepo) BookService {
	return &bookService{
		repo: repo,
	}
}

// CreateBook implements BookService.
func (s *bookService) CreateBook(ctx context.Context, book *model.Book) (*model.Book, error) {
	return s.repo.CreateBook(ctx, book)
}

// GetBook implements BookService.
func (s *bookService) GetBookById(ctx context.Context, id string) (*model.Book, error) {
	return s.repo.GetBookById(ctx, id)
}

// ListBooks implements BookService.
func (s *bookService) ListBooks(ctx context.Context) ([]*model.Book, error) {
	books, err := s.repo.GetBooks(ctx, nil)
	if err != nil {
		return nil, err
	}

	var bookPointers []*model.Book
	for _, book := range books {
		bookPointers = append(bookPointers, &book)
	}

	return bookPointers, nil
}
