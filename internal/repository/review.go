package repository

import (
	"context"
	"fmt"
	"github.com/forkpoons/reviews-sersice/internal/dto"
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

func (r *Review) Add(ctx context.Context, review dto.Review) error {
	q := `INSERT INTO promo (id, promo_type, entity, available_categories, discount, discount_type, expired_at, entity_id) VALUES (@id, @promo_type, @entity, @available_categories, @discount, @discount_type, @expired_at, @entity_id)`
	args := pgx.NamedArgs{
		"id":                   promo.ID,
		"promo_type":           promo.PromoType,
		"entity":               promo.Entity,
		"available_categories": promo.AvailableCategories,
		"discount":             promo.Discount,

		"expired_at": promo.ExpiredAt,
		"entity_id":  promo.EntityID,
	}
	_, err := r.pool.Exec(ctx, q, args)
	if err != nil {
		return fmt.Errorf("unable to insert row: %w", err)
	}
	return nil
}
