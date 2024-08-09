package service

import (
	"context"

	"github.com/storyofhis/auth-svc/internal/repository/gorm"
	"github.com/storyofhis/auth-svc/internal/repository/model"
)

type borrowService struct {
	repo *gorm.BorrowRepo
}

// CreateBorrow implements BorrowService.
func (s *borrowService) CreateBorrow(ctx context.Context, borrow *model.Borrow) (*model.Borrow, error) {
	return s.repo.CreateBorrow(ctx, borrow)
}

// GetBorrowById implements BorrowService.
func (s *borrowService) GetBorrowById(ctx context.Context, id string) (*model.Borrow, error) {
	return s.repo.GetBorrowById(ctx, id)
}

// GetBorrows implements BorrowService.
func (s *borrowService) GetBorrows(ctx context.Context) ([]*model.Borrow, error) {
	borrows, err := s.repo.GetBorrows(ctx, nil)
	if err != nil {
		return nil, err
	}

	var borrowPointers []*model.Borrow
	for _, borrow := range borrows {
		borrowPointers = append(borrowPointers, borrow)
	}

	return borrowPointers, nil
}

func NewBorrowService(repo *gorm.BorrowRepo) BorrowService {
	return &borrowService{
		repo: repo,
	}
}

type BorrowService interface {
	CreateBorrow(ctx context.Context, borrow *model.Borrow) (*model.Borrow, error)
	GetBorrowById(ctx context.Context, id string) (*model.Borrow, error)
	GetBorrows(ctx context.Context) ([]*model.Borrow, error)
}
