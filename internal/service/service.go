package service

import (
	"BDServer/internal/models"
	"BDServer/internal/repository"
	"context"
	"time"
)

type UserService interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
	ListUsers(ctx context.Context, c models.Conditions) ([]models.User, error)
}

type UserServ struct {
	repo repository.UserRepository
}

func New(repo repository.UserRepository) *UserServ {
	return &UserServ{
		repo: repo,
	}
}

func (s *UserServ) CreateUser(ctx context.Context, u *models.User) error {
	return s.repo.Create(ctx, *u)
}

func (s *UserServ) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	r, err := s.repo.GetByID(ctx, id)
	return &r, err
}

func (s *UserServ) UpdateUser(ctx context.Context, u *models.User) error {
	u.UpdateAt = time.Now()
	return s.repo.Update(ctx, *u)
}

func (s *UserServ) DeleteUser(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *UserServ) ListUsers(ctx context.Context, c models.Conditions) ([]models.User, error) {
	return s.repo.List(ctx, c)
}
