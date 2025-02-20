package dto

type Review struct {
	ReviewType string  `json:"review_type"`
	ReviewText string  `json:"review_text"`
	Media      []byte  `json:"media"`
	ProductId  string  `json:"product_id"`
	Rate       float64 `json:"rate"`
}
