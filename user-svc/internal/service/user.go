package service

import (
	"context"

	"github.com/storyofhis/auth-svc/internal/repository/gorm"
	"github.com/storyofhis/auth-svc/internal/repository/model"
)

type userService struct {
	repo *gorm.UserRepo
}

func NewUserService(repo *gorm.UserRepo) UserService {
	return &userService{
		repo: repo,
	}
}

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserById(ctx context.Context, id string) (*model.User, error)
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

func (s *userService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return s.repo.CreateUser(ctx, user)
}

func (s *userService) GetUserById(ctx context.Context, id string) (*model.User, error) {
	return s.repo.GetUserById(ctx, id)
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *userService) GetUsers(ctx context.Context) ([]*model.User, error) {
	users, err := s.repo.GetUsers(ctx, nil)
	if err != nil {
		return nil, err
	}

	var userPointers []*model.User
	for _, user := range users {
		userPointers = append(userPointers, user)
	}
	return userPointers, nil
}
