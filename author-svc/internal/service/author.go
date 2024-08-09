package service

import (
	"context"

	"github.com/storyofhis/author-service/internal/repository/gorm"
	"github.com/storyofhis/author-service/internal/repository/model"
)

type authorService struct {
	repo *gorm.AuthorRepo
}

// CreateAuthor implements AuthorService.
func (s *authorService) CreateAuthor(ctx context.Context, author *model.Author) (*model.Author, error) {
	return s.repo.CreateAuthor(ctx, author)
}

// GetAuthorById implements AuthorService.
func (s *authorService) GetAuthorById(ctx context.Context, id string) (*model.Author, error) {
	return s.repo.GetAuthorById(ctx, id)
}

// GetAuthors implements AuthorService.
func (s *authorService) GetAuthors(ctx context.Context) ([]*model.Author, error) {
	authors, err := s.repo.GetAuthors(ctx, nil)
	if err != nil {
		return nil, err
	}

	var authorPointers []*model.Author
	for _, author := range authors {
		authorPointers = append(authorPointers, author)
	}
	return authorPointers, nil
}

type AuthorService interface {
	CreateAuthor(ctx context.Context, author *model.Author) (*model.Author, error)
	GetAuthorById(ctx context.Context, id string) (*model.Author, error)
	GetAuthors(ctx context.Context) ([]*model.Author, error)
}

func NewAuthorService(repo *gorm.AuthorRepo) AuthorService {
	return &authorService{
		repo: repo,
	}
}
