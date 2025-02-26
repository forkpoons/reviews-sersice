package dto

import (
	"github.com/google/uuid"
	"time"
)

type ReviewDB struct {
	ID         uuid.UUID `json:"id"`
	Rtype      string    `json:"rtype"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ProductId  uuid.UUID `json:"product_id"`
	UserID     uuid.UUID `json:"user_id"`
	ReviewText string    `json:"review_text"`
	Media      string    `json:"media"`
	Rate       int       `json:"rate"`
	Status     string    `json:"status"`
}

type Review struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ProductId  uuid.UUID `json:"product_id"`
	UserID     uuid.UUID `json:"user_id"`
	ReviewText string    `json:"review_text"`
	Media      string    `json:"media"`
	Rate       int       `json:"rate"`
	Status     string    `json:"status"`
}

type Question struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ProductId    uuid.UUID `json:"product_id"`
	UserID       uuid.UUID `json:"user_id"`
	QuestionText string    `json:"review_text"`
	Media        string    `json:"media"`
	Status       string    `json:"status"`
	Answers      []Answer
}

type Answer struct {
	ID         uuid.UUID `json:"id"`
	AnswerText string    `json:"answer_text"`
	UserID     uuid.UUID `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	QuestionID uuid.UUID `json:"question_id"`
	Status     string    `json:"status"`
}
