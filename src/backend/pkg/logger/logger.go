package logger

import (
	"log/slog"
	"os"

	"github.com/jinpain/go-auth-service/pkg/configloader"
)

func New(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case configloader.EnvLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case configloader.EnvProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
