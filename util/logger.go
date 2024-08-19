package util

import (
	"go.uber.org/zap"
	"log"
)

func NewLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to create dwmbar logger: %s\n", err)
	}

	return logger
}
