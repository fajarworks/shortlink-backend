package repository

import (
	"context"
	"errors"

	"github.com/fajarworks/shortlink-backend/internal/model"
	"github.com/fajarworks/shortlink-backend/internal/utils"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LinkRepo struct {
	db *pgxpool.Pool
}

func NewLinkRepo(db *pgxpool.Pool) *LinkRepo {
	return &LinkRepo{
		db: db,
	}
}

func (r *LinkRepo) CreateLink(ctx context.Context, link model.Link) (model.Link, error) {
	query := `
		INSERT INTO shortlinks (user_id, slug, destination_url)
		VALUES ($1, $2, $3)
		RETURNING id, created_at`

	var data model.Link
	err := r.db.QueryRow(ctx, query, &link.UserId, &link.Slug, &link.OriginalURL).
		Scan(&data.Id, &data.CreatedAt)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return model.Link{}, utils.ERrSlugTaken
		}
		return model.Link{}, err
	}
	return data, nil
}
