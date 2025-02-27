package adminAPI

import (
	"context"
	"encoding/json"
	"github.com/fasthttp/router"
	"github.com/forkpoons/reviews-sersice/internal/dto"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"net/http"
)

type reviewRepo interface {
	GetReviewsByStatus(ctx context.Context, productID uuid.UUID, status []string) (*[]dto.Review, error)
	GetQuestionsByStatus(ctx context.Context, productID uuid.UUID, status []string) (*[]dto.Question, error)
	AddReview(ctx context.Context, review dto.Review) error
	EditReview(ctx context.Context, review dto.Review) error
	GetProductRate(ctx context.Context, productID uuid.UUID) (float32, error)
	SetReviewStatus(ctx context.Context, id uuid.UUID, status string) error
	SetAnswerStatus(ctx context.Context, id uuid.UUID, status string) error
}

type Service struct {
	r          *router.Router
	log        zerolog.Logger
	reviewRepo reviewRepo
}

func NewService(log zerolog.Logger, reviewRepo reviewRepo, appSecret string) *Service {
	r := router.New()
	s := Service{
		log:        log,
		reviewRepo: reviewRepo,
	}
	r.GET("/api/reviews", s.GetReviews)
	r.GET("/api/questions", s.GetQuestion)
	r.PUT("/api/review", s.SetReviewStatus)
	r.PUT("/api/questions", s.SetQuestionStatus)
	r.PUT("/api/answer", s.SetAnswerStatus)
	s.r = r
	return &s
}

func (s *Service) Start(ctx context.Context) error {
	server := fasthttp.Server{
		Handler: s.r.Handler,
		Name:    "Review admin API",
	}
	emergencyShutdown := make(chan error)
	go func() {
		s.log.Info().Msgf("Starting admin server")
		err := server.ListenAndServe(":9090")
		emergencyShutdown <- err
	}()

	select {
	case <-ctx.Done():
		return server.Shutdown()
	case e := <-emergencyShutdown:
		return e
	}
}

func (s *Service) GetReviews(ctx *fasthttp.RequestCtx) {
	productID, err := uuid.ParseBytes(ctx.QueryArgs().Peek("product_id"))
	if err != nil {

	}
	review, err := s.reviewRepo.GetReviewsByStatus(ctx, productID, []string{"created", "published"})

	ctx.SetContentType("application/json")
	data, err := json.Marshal(review)
	if err != nil {
		s.log.Debug().Err(err).Msg("Error marshaling reviews")
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(data)
}

func (s *Service) GetQuestion(ctx *fasthttp.RequestCtx) {
	productID, err := uuid.ParseBytes(ctx.QueryArgs().Peek("product_id"))
	if err != nil {

	}
	review, err := s.reviewRepo.GetQuestionsByStatus(ctx, productID, []string{"created", "published"})
	ctx.SetContentType("application/json")
	data, err := json.Marshal(review)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(data)
}

func (s *Service) SetReviewStatus(ctx *fasthttp.RequestCtx) {
	ID, err := uuid.ParseBytes(ctx.QueryArgs().Peek("id"))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	status := ctx.QueryArgs().Peek("status")
	err = s.reviewRepo.SetReviewStatus(ctx, ID, string(status))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.Response.SetStatusCode(http.StatusOK)
}

func (s *Service) SetQuestionStatus(ctx *fasthttp.RequestCtx) {
	ID, err := uuid.ParseBytes(ctx.QueryArgs().Peek("id"))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	status := ctx.QueryArgs().Peek("status")
	err = s.reviewRepo.SetReviewStatus(ctx, ID, string(status))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.Response.SetStatusCode(http.StatusOK)
}

func (s *Service) SetAnswerStatus(ctx *fasthttp.RequestCtx) {
	ID, err := uuid.ParseBytes(ctx.QueryArgs().Peek("id"))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	status := ctx.QueryArgs().Peek("status")
	err = s.reviewRepo.SetAnswerStatus(ctx, ID, string(status))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.Response.SetStatusCode(http.StatusOK)
}
