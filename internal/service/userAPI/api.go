package userAPI

import (
	"context"
	"github.com/fasthttp/router"
	"github.com/forkpoons/reviews-sersice/internal/dto"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
)

type reviewRepo interface {
	GetReviews(ctx context.Context, id uuid.UUID) (*[]dto.Review, error)
	AddReview(ctx context.Context, review dto.Review) error
	EditReview(ctx context.Context, review dto.Review) error
	DeleteReview(ctx context.Context, id uuid.UUID) error
	GetProductRate(ctx context.Context, productID uuid.UUID) (float32, error)
}

type Service struct {
	r          *router.Router
	log        zerolog.Logger
	reviewRepo reviewRepo
	appSecret  string
}

func NewService(log zerolog.Logger, reviewRepo reviewRepo, appSecret string) *Service {
	r := router.New()
	s := Service{
		log:        log,
		reviewRepo: reviewRepo,
		appSecret:  appSecret,
	}
	r.POST("/api/review", s.AddReview)
	r.PUT("/api/review", s.EditReview)
	r.DELETE("/api/review", s.DeleteReview)
	r.GET("/api/reviews", s.GetReviews)
	r.GET("/api/review", s.GetReviewByID)
	r.GET("/api/rate", s.GetRate)

	s.r = r
	return &s
}

func (s *Service) Start(ctx context.Context) error {
	server := fasthttp.Server{
		Handler: s.r.Handler,
		Name:    "Promo API",
	}
	emergencyShutdown := make(chan error)
	go func() {
		s.log.Info().Msgf("Starting server21321")
		err := server.ListenAndServe(":8080")
		emergencyShutdown <- err
	}()

	select {
	case <-ctx.Done():
		return server.Shutdown()
	case e := <-emergencyShutdown:
		return e
	}
}
