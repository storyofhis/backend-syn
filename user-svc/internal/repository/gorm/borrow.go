package gorm

import (
	"context"

	"github.com/storyofhis/auth-svc/internal/repository/model"
	"gorm.io/gorm"
)

type BorrowRepo struct {
	db *gorm.DB
}

func NewBorrowRepo(db *gorm.DB) *BorrowRepo {
	return &BorrowRepo{
		db: db,
	}
}

func (repo *BorrowRepo) CreateBorrow(ctx context.Context, borrow *model.Borrow) (*model.Borrow, error) {
	if err := repo.db.WithContext(ctx).Create(borrow).Error; err != nil {
		return nil, err
	}
	return borrow, nil
}

func (repo *BorrowRepo) GetBorrowById(ctx context.Context, id string) (*model.Borrow, error) {
	borrow := new(model.Borrow)
	if err := repo.db.WithContext(ctx).Where("id = ?", id).Take(borrow).Error; err != nil {
		return nil, err
	}
	return borrow, nil
}

func (repo *BorrowRepo) GetBorrows(ctx context.Context, criteria map[string]interface{}) ([]*model.Borrow, error) {
	var borrows []*model.Borrow
	query := repo.db.WithContext(ctx).Model(&model.Borrow{})

	for key, value := range criteria {
		query = query.Where(key+" = ?", value)
	}

	if err := query.Find(&borrows).Error; err != nil {
		return nil, err
	}
	return borrows, nil
}
