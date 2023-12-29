package logger

import (
	"log/slog"
	"os"
)

func Init() {
	lgr := slog.New(
		slog.NewJSONHandler(os.Stdout,
			&slog.HandlerOptions{
				AddSource: true,
			},
		))

	slog.SetDefault(lgr)
}
