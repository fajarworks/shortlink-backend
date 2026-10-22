package repository

import (
	"context"
	"errors"

	"github.com/fajarworks/shortlink-backend/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepo struct {
	db *pgxpool.Pool
}

func NewAuthRepo(db *pgxpool.Pool) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) CreateUser(ctx context.Context, user model.User) error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2)`
	cmdTag, err := r.db.Exec(ctx, query, user.Email, user.Password)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return errors.New("no rows affected")
	}
	return err

}

func (r *AuthRepo) FindUserByEmail(ctx context.Context, email string) (model.User, error) {
	query := `SELECT id, email,  COALESCE(fullname, '') AS fullname, 
    COALESCE(image_url, '') AS image_url, password FROM users WHERE email = $1`
	var user model.User
	if err := r.db.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.FullName, &user.ImageURL, &user.Password); err != nil {
		return model.User{}, err
	}
	return user, nil
}
