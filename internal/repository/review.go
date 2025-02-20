package repository

import (
	"context"
	"github.com/forkpoons/reviews-sersice/internal/dto"
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

func (r *Review) Add(ctx context.Context, promo dto.Review) error {
	return nil
}
