package entity

type IncrementDTO struct {
	Key   string `json:"key" validate:"required"`
	Value *int64 `json:"value" validate:"required"`
}

type IncrementResult struct {
	Key   string
	Value int64 `redis:"number"`
}
