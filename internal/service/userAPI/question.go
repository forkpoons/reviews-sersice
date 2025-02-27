package userAPI

import (
	"bytes"
	"encoding/json"
	"github.com/forkpoons/reviews-sersice/internal/dto"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"net/http"
)

func (s *Service) GetQuestionByID(ctx *fasthttp.RequestCtx) {
	productID, err := uuid.ParseBytes(ctx.QueryArgs().Peek("product_id"))
	if err != nil {

	}
	review, err := s.reviewRepo.GetReviewsByStatus(ctx, productID, []string{"created", "published"})
	ctx.SetContentType("application/json")
	data, err := json.Marshal(review)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(data)
}

func (s *Service) GetQuestions(ctx *fasthttp.RequestCtx) {
	productID, err := uuid.ParseBytes(ctx.QueryArgs().Peek("product_id"))
	if err != nil {

	}
	review, err := s.reviewRepo.GetReviewsByStatus(ctx, productID, []string{"created", "published"})
	ctx.SetContentType("application/json")
	data, err := json.Marshal(review)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(data)
}

func (s *Service) AddQuestion(ctx *fasthttp.RequestCtx) {
	decoder := json.NewDecoder(bytes.NewReader(ctx.PostBody()))
	s.log.Info().Msgf("Received request to add review")
	review := dto.Question{}
	if err := decoder.Decode(&review); err != nil {
		s.log.Error().Err(err).Send()
		ctx.Response.SetStatusCode(http.StatusBadRequest)
		return
	}
	var err error
	review.UserID, err = uuid.Parse(ctx.UserValue("uid").(string))
	if err != nil {
		s.log.Error().Err(err).Send()
		ctx.Response.SetStatusCode(http.StatusBadRequest)
		return
	}
	err = s.reviewRepo.AddQuestion(ctx, review)
	if err != nil {
		return
	}

	ctx.Response.SetStatusCode(http.StatusCreated)
}

func (s *Service) EditQuestion(ctx *fasthttp.RequestCtx) {
	decoder := json.NewDecoder(bytes.NewReader(ctx.PostBody()))
	s.log.Info().Msgf("Received request to add review")
	review := dto.Question{}
	if err := decoder.Decode(&review); err != nil {
		s.log.Error().Err(err).Send()
		ctx.Response.SetStatusCode(http.StatusBadRequest)
		return
	}

	err := s.reviewRepo.EditQuestion(ctx, review)
	if err != nil {
		return
	}

	ctx.Response.SetStatusCode(http.StatusCreated)
}

func (s *Service) DeleteQuestion(ctx *fasthttp.RequestCtx) {
	productID, err := uuid.ParseBytes(ctx.QueryArgs().Peek("id"))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	err = s.reviewRepo.SetReviewStatus(ctx, productID, "deleted")
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.Response.SetStatusCode(http.StatusCreated)
}

func (s *Service) GetQuestion(ctx *fasthttp.RequestCtx) {
	ID, err := uuid.ParseBytes(ctx.QueryArgs().Peek("id"))
	if err != nil {

	}
	review, err := s.reviewRepo.GetQuestionsByStatus(ctx, ID, []string{"published"})
	ctx.SetContentType("application/json")
	data, err := json.Marshal(review)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(data)
}

func (s *Service) AddAnswer(ctx *fasthttp.RequestCtx) {
	decoder := json.NewDecoder(bytes.NewReader(ctx.PostBody()))
	s.log.Info().Msgf("Received request to add review")
	answer := dto.Answer{}
	if err := decoder.Decode(&answer); err != nil {
		s.log.Error().Err(err).Send()
		ctx.Response.SetStatusCode(http.StatusBadRequest)
		return
	}
	var err error
	answer.UserID, err = uuid.Parse(ctx.UserValue("uid").(string))
	if err != nil {
		s.log.Error().Err(err).Send()
		ctx.Response.SetStatusCode(http.StatusBadRequest)
		return
	}
	err = s.reviewRepo.AddAnswer(ctx, answer)
	if err != nil {
		return
	}

	ctx.Response.SetStatusCode(http.StatusCreated)
}

func (s *Service) EditAnswer(ctx *fasthttp.RequestCtx) {
	decoder := json.NewDecoder(bytes.NewReader(ctx.PostBody()))
	s.log.Info().Msgf("Received request to add review")
	review := dto.Review{}
	if err := decoder.Decode(&review); err != nil {
		s.log.Error().Err(err).Send()
		ctx.Response.SetStatusCode(http.StatusBadRequest)
		return
	}

	err := s.reviewRepo.EditReview(ctx, review)
	if err != nil {
		return
	}

	ctx.Response.SetStatusCode(http.StatusCreated)
}
