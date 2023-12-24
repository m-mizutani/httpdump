package utils

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func init() {
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))
}

func Logger() *slog.Logger {
	return logger
}

func ReplaceLogger(l *slog.Logger) {
	logger = l
}
