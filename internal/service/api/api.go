package api

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/fasthttp/router"
	"github.com/forkpoons/reviews-sersice/internal/dto"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"net/http"
)

type reviewRepo interface {
	Add(ctx context.Context, review dto.Review) error
}

type service struct {
	r          *router.Router
	log        zerolog.Logger
	reviewRepo reviewRepo
	appSecret  string
}

func NewService(log zerolog.Logger, reviewRepo reviewRepo, appSecret string) *service {
	r := router.New()
	s := service{
		log:        log,
		reviewRepo: reviewRepo,
		appSecret:  appSecret,
	}
	r.POST("/api/review", s.AddReview)

	s.r = r
	return &s
}

func (s *service) Start(ctx context.Context) error {
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

func (s *service) AddReview(ctx *fasthttp.RequestCtx) {
	decoder := json.NewDecoder(bytes.NewReader(ctx.PostBody()))
	s.log.Info().Msgf("Received request to add review")
	review := dto.Review{
		ReviewType: "review",
	}
	if err := decoder.Decode(&review); err != nil {
		s.log.Error().Err(err).Send()
		ctx.Response.SetStatusCode(http.StatusBadRequest)
		return
	}

	err := s.reviewRepo.Add(ctx, review)
	if err != nil {
		return
	}

	ctx.Response.SetStatusCode(http.StatusCreated)
}
