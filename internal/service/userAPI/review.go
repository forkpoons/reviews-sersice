package userAPI

import (
	"bytes"
	"encoding/json"
	"github.com/forkpoons/reviews-sersice/internal/dto"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"net/http"
)

func (s *Service) GetReviewByID(ctx *fasthttp.RequestCtx) {
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

func (s *Service) GetReviews(ctx *fasthttp.RequestCtx) {
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

func (s *Service) GetRate(ctx *fasthttp.RequestCtx) {
	productID, err := uuid.ParseBytes(ctx.QueryArgs().Peek("product_id"))
	if err != nil {

	}
	rate, err := s.reviewRepo.GetProductRate(ctx, productID)
	ctx.SetContentType("application/json")
	data, err := json.Marshal(rate)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(data)
}

func (s *Service) AddReview(ctx *fasthttp.RequestCtx) {
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

func (s *Service) EditReview(ctx *fasthttp.RequestCtx) {
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

func (s *Service) DeleteReview(ctx *fasthttp.RequestCtx) {
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
