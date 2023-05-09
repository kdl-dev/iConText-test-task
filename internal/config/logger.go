package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

var (
	loggerConfigFilePath = "config/logger.yaml"
	logger               Logger
)

func init() {
	if err := cleanenv.ReadConfig(loggerConfigFilePath, &logger); err != nil {
		log.Fatal(err)
	}

	AppCfg.Logger = logger
}

type Logger struct {
	Outputs []struct {
		Storage  string `yaml:"storage" validate:"required"`
		Level    string `yaml:"level" validate:"required"`
		Encoding string `yaml:"encoding" validate:"required"`
	} `yaml:"outputs" validate:"dive"`
	AddCaller       bool   `yaml:"add_caller" validate:"required"`
	StackTraceLevel string `yaml:"stacktrace_level" validate:"required"`
}
