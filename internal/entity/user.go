package entity

type UserDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name" validate:"required"`
	Age  *int64 `json:"age" validate:"required,gte=0"`
}

type User struct {
	ID   int64  `db:"user_id"`
	Name string `db:"name"`
	Age  int64  `db:"age"`
}
