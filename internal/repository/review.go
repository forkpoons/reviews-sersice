package repository

import (
	"context"
	"fmt"
	"github.com/forkpoons/reviews-sersice/internal/dto"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"time"
)

type Review struct {
	ctx  context.Context
	pool *pgxpool.Pool
	log  zerolog.Logger
}

func NewPromo(ctx context.Context, pool *pgxpool.Pool, log zerolog.Logger) *Review {
	return &Review{ctx: ctx, pool: pool, log: log}
}

func (r *Review) GetReviewByID(ctx context.Context, ID uuid.UUID) (*dto.ReviewDB, error) {
	q := `SELECT * FROM reviews WHERE id = $1`

	rows, err := r.pool.Query(ctx, q, ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pr, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dto.ReviewDB])
	if err != nil {
		return nil, err
	}

	return &pr, nil
}

func (r *Review) GetReviews(ctx context.Context, productID uuid.UUID) (*[]dto.ReviewDB, error) {
	q := `SELECT * FROM reviews WHERE product_id = $1 AND rtype = 'review'`
	rows, err := r.pool.Query(ctx, q, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rdb, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.ReviewDB])

	if err != nil {
		r.log.Debug().Err(err).Send()
		return nil, err
	}
	return &rdb, nil
}

func (r *Review) GetProductRate(ctx context.Context, productID uuid.UUID) (float32, error) {
	q := `SELECT rate FROM reviews WHERE product_id = $1 AND rtype = 'review' AND status = 'published'`
	rows, err := r.pool.Query(ctx, q, productID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var sum float32
	var count int
	for rows.Next() {
		var rate float32
		if err := rows.Scan(&rate); err != nil {
			return 0, err
		}
		sum += rate
		count++
	}
	if err := rows.Err(); err != nil {
		return 0, err
	}

	if count == 0 {
		return 0, nil
	}
	average := sum / float32(count)
	return average, nil
}

func (r *Review) AddReview(ctx context.Context, review dto.Review) error {
	q := `INSERT INTO reviews (id, rtype, created_at, updated_at, product_id, user_id, review_text, media, rate, status) VALUES (@id, @type, @created_at, @updated_at, @product_id, @user_id, @review_text, @media, @rate, @status)`
	args := pgx.NamedArgs{
		"id":          uuid.New(),
		"type":        "review",
		"created_at":  time.Now(),
		"updated_at":  time.Now(),
		"product_id":  review.ProductId,
		"user_id":     review.UserID,
		"review_text": review.ReviewText,
		"media":       review.Media,
		"rate":        review.Rate,
		"status":      "created",
	}
	_, err := r.pool.Exec(ctx, q, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}
	return nil
}

func (r *Review) EditReview(ctx context.Context, review dto.Review) error {
	q := `UPDATE reviews SET updated_at = @updated_at, review_text = @review_text, media = @media, rate = @rate WHERE id = @id;`
	args := pgx.NamedArgs{
		"id":          review.ID,
		"review_text": review.ReviewText,
		"updated_at":  time.Now(),
		"media":       review.Media,
		"rate":        review.Rate,
	}
	_, err := r.pool.Exec(ctx, q, args)
	if err != nil {
		return err
	}
	return nil
}

func (r *Review) DeleteReview(ctx context.Context, id uuid.UUID) error {
	q := `UPDATE reviews SET status = 'deleted' WHERE id = $1`

	_, err := r.pool.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	return nil
}
