package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	envConfigFilePath = ".env"
	AppCfg            App
)

func init() {
	if err := cleanenv.ReadConfig(envConfigFilePath, &AppCfg); err != nil {
		log.Fatal(err)
	}

	AppCfg.Logger = logger
}

type App struct {
	Server   `validate:"dive"`
	Postgres `validate:"dive"`
	Redis    `validate:"dive"`
	Logger   `validate:"dive"`
}
