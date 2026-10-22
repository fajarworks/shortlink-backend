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

func (r *LinkRepo) FindLinksByUserID(ctx context.Context, userID int) ([]model.Link, error) {
	query := `
        SELECT
            id,
            user_id,
            destination_url,
            slug,
            click_count,
            created_at,
            deleted_at
        FROM shortlinks
        WHERE user_id = $1
          AND deleted_at IS NULL
        ORDER BY created_at DESC
    `

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	links := make([]model.Link, 0)

	for rows.Next() {
		var link model.Link

		err := rows.Scan(
			&link.Id,
			&link.UserId,
			&link.OriginalURL,
			&link.Slug,
			&link.ClickCount,
			&link.CreatedAt,
			&link.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return links, nil
}
