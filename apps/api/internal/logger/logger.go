package logger

import (
	"log/slog"
	"os"

	"github.com/dtt4h/devodoro/internal/config"
)

func New(cfg config.Logger) *slog.Logger {
	var level slog.Level
	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handler).With("service", "devodoro")

	return logger
}
