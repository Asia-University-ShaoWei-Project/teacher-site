package log

import (
	"context"

	"go.uber.org/zap"
)

func NewLogzap(ctx context.Context) *zap.Logger {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{
		"./record.log",
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger
}
