package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func New() zerolog.Logger {
	// Pretty console output for dev
	cw := zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = time.RFC3339
	})
	cw.Out = os.Stdout
	log := zerolog.New(cw).With().Timestamp().Logger()
	return log
}
