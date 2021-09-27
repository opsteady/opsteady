package tasks

import (
	"github.com/rs/zerolog"
)

type LogWriter struct {
	logger *zerolog.Logger
}

func NewLogWriter(logger *zerolog.Logger) *LogWriter {
	lw := &LogWriter{}
	lw.logger = logger
	return lw
}

func (lw LogWriter) Write(p []byte) (n int, err error) {
	return lw.logger.Write(p)
}
