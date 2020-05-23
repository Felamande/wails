package runtime

import "github.com/Felamande/wails/lib/logger"

// Log exposes the logging interface to the runtime
type Log struct{}

// NewLog creates a new Log struct
func NewLog() *Log {
	return &Log{}
}

// New creates a new logger
func (r *Log) New(prefix string) *logger.CustomLogger {
	return logger.NewCustomLogger(prefix)
}
