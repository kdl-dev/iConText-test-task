package config

var (
	AppCfg *App
)

type App struct {
	Server   `validate:"dive"`
	Postgres `validate:"dive"`
	Redis    `validate:"dive"`
	Logger   `validate:"dive"`
}
