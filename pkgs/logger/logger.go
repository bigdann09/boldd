package logger

import "go.uber.org/zap"

func NewLogger(environment string) *zap.Logger {
	logger := zap.Must(zap.NewProduction())
	if environment == "development" {
		logger = zap.Must(zap.NewDevelopment())
	}
	return logger
}
