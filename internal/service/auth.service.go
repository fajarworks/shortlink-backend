package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/fajarworks/shortlink-backend/internal/dto"
	"github.com/fajarworks/shortlink-backend/internal/model"
	"github.com/fajarworks/shortlink-backend/internal/repository"
	"github.com/fajarworks/shortlink-backend/internal/utils"
	"github.com/fajarworks/shortlink-backend/pkg"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	repo *repository.AuthRepo
	rdb  *redis.Client
}

func NewAuthService(repo *repository.AuthRepo, rdb *redis.Client) *AuthService {
	return &AuthService{
		repo: repo,
		rdb:  rdb,
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

func (s *AuthService) Login(ctx context.Context, payload dto.LoginRequest) (dto.User, string, error) {
	if payload.Email == "" || len(payload.Password) == 0 {
		return dto.User{}, "", utils.ErrFieldEmpty
	}
	user, err := s.repo.FindUserByEmail(ctx, payload.Email)
	if err != nil {
		log.Println(err.Error())
		if errors.Is(err, pgx.ErrNoRows) {

			return dto.User{}, "", utils.ErrWrongEmailPass
		}
		return dto.User{}, "", err
	}
	if err := pkg.ComparePassAndHash(payload.Password, user.Password); err != nil {
		log.Println(err.Error())
		return dto.User{}, "", utils.ErrWrongEmailPass
	}

	claims := pkg.NewJWTClaims(user.ID, user.Email)
	token, err := claims.GenToken()
	if err != nil {
		return dto.User{}, "", err
	}
	return dto.User{
		Email: user.Email,
	}, token, nil

}

func (s *AuthService) Logout(ctx context.Context, userId int, jti string, expiresAt time.Time) error {

	key := fmt.Sprintf("shortlink:blacklist:%s", jti)
	ttl := time.Until(expiresAt)
	if ttl <= 0 {
		return nil
	}

	if err := s.rdb.Set(ctx, key, userId, ttl).Err(); err != nil {
		return err
	}
	return nil
}
