package config

type App struct {
	Server   `validate:"dive"`
	Postgres `validate:"dive"`
	Redis    `validate:"dive"`
	Logger   `validate:"dive"`
}
