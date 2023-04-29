package logger

import (
	"errors"
	"fmt"
	"os"

	"github.com/kdl-dev/iConText-test-task/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zap.Logger
}

func NewLogger(cfg *config.Logger) (*Logger, error) {

	cores, err := getLogCores(cfg, zapcore.ISO8601TimeEncoder)
	if err != nil {
		return nil, fmt.Errorf("get log cores error: %w", err)
	}

	core := zapcore.NewTee(cores...)

	var options []zap.Option

	if cfg.AddCaller {
		options = append(options, zap.AddCaller())
	}

	stackTraceLevel, err := zap.ParseAtomicLevel(cfg.StackTraceLevel)
	if err != nil {
		return nil, fmt.Errorf("stacktrace level parse error: %w", err)
	}

	options = append(options, zap.AddStacktrace(stackTraceLevel))

	logger := zap.New(core, options...)

	return &Logger{*logger}, nil
}

func getLogCores(cfg *config.Logger, timeEncoder zapcore.TimeEncoder) ([]zapcore.Core, error) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = timeEncoder

	var cores []zapcore.Core

	for _, val := range cfg.Outputs {

		file, err := os.OpenFile(val.Storage, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, fmt.Errorf("file open error: %w", err)
		}

		encoder, err := parseEncoder(val.Encoding, config)
		if err != nil {
			return nil, fmt.Errorf("encoder type parse error: %w", err)
		}

		logWriter := zapcore.AddSync(file)

		logLevel, err := zap.ParseAtomicLevel(val.Level)
		if err != nil {
			return nil, fmt.Errorf("log level parse error: %w", err)
		}

		cores = append(cores, zapcore.NewCore(encoder, logWriter, logLevel))
	}

	return cores, nil
}

func parseEncoder(typeEncoder string, cfg zapcore.EncoderConfig) (zapcore.Encoder, error) {
	switch typeEncoder {
	case "console":
		return zapcore.NewConsoleEncoder(cfg), nil
	case "json":
		return zapcore.NewJSONEncoder(cfg), nil
	default:
		return nil, errors.New("unknown type encoder")
	}
}
