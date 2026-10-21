package service

import (
	"context"
	"errors"

	"github.com/fajarworks/shortlink-backend/internal/dto"
	"github.com/fajarworks/shortlink-backend/internal/model"
	"github.com/fajarworks/shortlink-backend/internal/repository"
	"github.com/fajarworks/shortlink-backend/internal/utils"
	"github.com/fajarworks/shortlink-backend/pkg"
	"github.com/jackc/pgx/v5"
)

type AuthService struct {
	repo *repository.AuthRepo
}

func NewAuthService(repo *repository.AuthRepo) *AuthService {
	return &AuthService{
		repo: repo,
	}

}

func (s *AuthService) CreateUser(ctx context.Context, payload dto.CreateUser) error {
	if payload.Email == "" || payload.Password == "" {
		return utils.ErrFieldEmpty
	}

	_, err := s.repo.FindUserByEmail(ctx, payload.Email)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	if err == nil {
		return utils.ErrEmailAlreadyExist
	}

	hashConf := pkg.NewRecomHashConfig()
	hashedPass, _ := hashConf.GenHash(payload.Password)

	user := model.User{
		Email:    payload.Email,
		Password: hashedPass,
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil

}
