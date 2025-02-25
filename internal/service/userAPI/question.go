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
	review, err := s.reviewRepo.GetReviews(ctx, productID)
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
	review, err := s.reviewRepo.GetReviews(ctx, productID)
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
	review := dto.Review{}
	if err := decoder.Decode(&review); err != nil {
		s.log.Error().Err(err).Send()
		ctx.Response.SetStatusCode(http.StatusBadRequest)
		return
	}

	err := s.reviewRepo.AddReview(ctx, review)
	if err != nil {
		return
	}

	ctx.Response.SetStatusCode(http.StatusCreated)
}

func (s *Service) EditQuestion(ctx *fasthttp.RequestCtx) {
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

func (s *Service) DeleteQuestion(ctx *fasthttp.RequestCtx) {
	productID, err := uuid.ParseBytes(ctx.QueryArgs().Peek("product_id"))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	err = s.reviewRepo.DeleteReview(ctx, productID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.Response.SetStatusCode(http.StatusCreated)
}

func (s *Service) GetAnswer(ctx *fasthttp.RequestCtx) {
	productID, err := uuid.ParseBytes(ctx.QueryArgs().Peek("product_id"))
	if err != nil {

	}
	review, err := s.reviewRepo.GetReviews(ctx, productID)
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
	review := dto.Review{}
	if err := decoder.Decode(&review); err != nil {
		s.log.Error().Err(err).Send()
		ctx.Response.SetStatusCode(http.StatusBadRequest)
		return
	}

	err := s.reviewRepo.AddReview(ctx, review)
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

func (s *Service) DeleteAnswer(ctx *fasthttp.RequestCtx) {
	productID, err := uuid.ParseBytes(ctx.QueryArgs().Peek("product_id"))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	err = s.reviewRepo.DeleteReview(ctx, productID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	ctx.Response.SetStatusCode(http.StatusCreated)
}
