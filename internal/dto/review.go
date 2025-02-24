package dto

import "github.com/google/uuid"

type Review struct {
	ReviewType string    `json:"review_type"`
	UserID     uuid.UUID `json:"user_id"`
	ReviewText string    `json:"review_text"`
	Media      []string  `json:"media"`
	ProductId  uuid.UUID `json:"product_id"`
	Rate       float32   `json:"rate"`
}
