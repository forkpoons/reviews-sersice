package repository

import (
	"context"
	"fmt"
	"github.com/forkpoons/reviews-sersice/internal/dto"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type Review struct {
	ctx  context.Context
	pool *pgxpool.Pool
	log  zerolog.Logger
}

func NewPromo(ctx context.Context, pool *pgxpool.Pool, log zerolog.Logger) *Review {
	return &Review{ctx: ctx, pool: pool, log: log}
}
func (r *Review) GetReview(ctx context.Context, id uuid.UUID) (*dto.Review, error) {
	q := `SELECT * FROM reviews WHERE id = $1`
	rows, err := r.pool.Query(ctx, q, id)
	if err != nil {
		r.log.Err(err).Send()
		return nil, err
	}
	defer rows.Close()

	pr, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[dto.Review])
	if err != nil {
		return nil, err
	}
	return &pr, nil
}

func (r *Review) AddReview(ctx context.Context, review dto.Review) error {
	q := `INSERT INTO reviews (id, user_id, review_text, media, product_id, rate) VALUES (@id, @user_id, @review_text, @media, @product_id, @rate)`
	args := pgx.NamedArgs{
		"id":          uuid.New(),
		"user_id":     review.UserID,
		"review_text": review.ReviewText,
		"media":       review.Media,
		"product_id":  review.ProductId,
		"rate":        review.Rate,
	}
	_, err := r.pool.Exec(ctx, q, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}
	return nil
}

func (r *Review) EditReview(ctx context.Context, review dto.Review) error {
	q := `INSERT INTO reviews (id, user_id, review_text, media, product_id, rate) VALUES (@id, @user_id, @review_text, @media, @product_id, @rate)`
	args := pgx.NamedArgs{
		"id":          uuid.New(),
		"user_id":     review.UserID,
		"review_text": review.ReviewText,
		"media":       review.Media,
		"product_id":  review.ProductId,
		"rate":        review.Rate,
	}
	_, err := r.pool.Exec(ctx, q, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}
	return nil
}

func (r *Review) DeleteReview(ctx context.Context, review dto.Review) error {
	q := `INSERT INTO reviews (id, user_id, review_text, media, product_id, rate) VALUES (@id, @user_id, @review_text, @media, @product_id, @rate)`
	args := pgx.NamedArgs{
		"id":          uuid.New(),
		"user_id":     review.UserID,
		"review_text": review.ReviewText,
		"media":       review.Media,
		"product_id":  review.ProductId,
		"rate":        review.Rate,
	}
	_, err := r.pool.Exec(ctx, q, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}
	return nil
}
