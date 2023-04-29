package entity

type UserDTO struct {
	Name string `json:"name" validate:"required"`
	Age  int64  `json:"age" validate:"required"`
}

type User struct {
	ID   uint64 `db:"user_id"`
	Name string `db:"name"`
	Age  int64  `db:"age"`
}
