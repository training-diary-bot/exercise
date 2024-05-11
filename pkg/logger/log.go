package logger

import (
	"github.com/rs/zerolog"
	"os"
)

func New() zerolog.Logger {
	return zerolog.New(os.Stderr)
}
