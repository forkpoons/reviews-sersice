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
	GetReviews(ctx context.Context, id uuid.UUID) (*[]dto.ReviewDB, error)
	GetReviewByID(ctx context.Context, ID uuid.UUID) (*dto.ReviewDB, error)
	AddReview(ctx context.Context, review dto.Review) error
	EditReview(ctx context.Context, review dto.Review) error
	DeleteReview(ctx context.Context, id uuid.UUID) error
	GetProductRate(ctx context.Context, productID uuid.UUID) (float32, error)
	GetQuestions(ctx context.Context, productID uuid.UUID) (*[]dto.Question, error)
	AddQuestion(ctx context.Context, review dto.Question) error
	EditQuestion(ctx context.Context, review dto.Question) error
	AddAnswer(ctx context.Context, answer dto.Answer) error
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
	r.POST("/api/review", s.auth(s.AddReview))
	r.PUT("/api/review", s.auth(s.EditReview))
	r.DELETE("/api/review", s.auth(s.DeleteReview))
	r.GET("/api/reviews", s.GetReviews)
	r.GET("/api/review", s.GetReviewByID)

	r.GET("/api/rate", s.GetRate)

	r.POST("/api/question", s.auth(s.AddQuestion))
	r.PUT("/api/question", s.auth(s.EditReview))
	r.DELETE("/api/question", s.auth(s.DeleteReview))
	r.GET("/api/questions", s.GetQuestion)
	r.GET("/api/question", s.GetReviewByID)

	r.POST("/api/answer", s.auth(s.AddAnswer))
	r.PUT("/api/answer", s.auth(s.EditReview))
	r.DELETE("/api/answer", s.auth(s.DeleteReview))

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
