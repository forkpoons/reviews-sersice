package api

import (
	"context"
	"github.com/fasthttp/router"
	"github.com/forkpoons/reviews-sersice/internal/dto"
	"github.com/rs/zerolog"
)

type reviewRepo interface {
	Add(ctx context.Context, promo dto.Review) error
}

type service struct {
	r          *router.Router
	log        zerolog.Logger
	reviewRepo reviewRepo
	appSecret  string
}

func NewService(log zerolog.Logger, reviewRepo reviewRepo, appSecret string) *service {
	s := service{
		log:        log,
		reviewRepo: reviewRepo,
		appSecret:  appSecret,
	}

	return &s
}

func (s *service) Start(ctx context.Context) error {
	err := s.reviewRepo.Add(ctx, dto.Review{})
	if err != nil {
		return err
	}
	return nil
}
