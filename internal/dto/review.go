package dto

import (
	"github.com/google/uuid"
	"time"
)

type Review struct {
	ID         uuid.UUID `json:"id"`
	Rtype      string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	ProductId  uuid.UUID `json:"product_id"`
	UserID     uuid.UUID `json:"user_id"`
	ReviewText string    `json:"review_text"`
	Media      string    `json:"media"`
	Rate       float32   `json:"rate"`
	Status     string    `json:"status"`
}
