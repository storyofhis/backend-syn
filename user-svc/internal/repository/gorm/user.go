package gorm

import (
	"context"
	"strings"

	"github.com/storyofhis/auth-svc/internal/repository/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := repo.db.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) GetUserById(ctx context.Context, id string) (*model.User, error) {
	user := new(model.User)
	if err := repo.db.WithContext(ctx).Where("id = ?", id).Take(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := new(model.User)
	if err := repo.db.WithContext(ctx).Where("LOWER(email) = ?", strings.ToLower(email)).Take(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) GetUsers(ctx context.Context, criteria map[string]interface{}) ([]*model.User, error) {
	var users []*model.User
	query := repo.db.WithContext(ctx).Model(&model.User{})

	for key, value := range criteria {
		query = query.Where(key+" = ?", value)
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
