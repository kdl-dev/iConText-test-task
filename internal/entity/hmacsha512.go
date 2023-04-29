package entity

type HMACSHA512DTO struct {
	Text string `json:"text" validate:"required"`
	Key  string `json:"key" validate:"required"`
}
