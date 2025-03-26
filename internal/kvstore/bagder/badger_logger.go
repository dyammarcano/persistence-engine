package bagder

import (
	"log/slog"
	"os"
)

// Logger is a simple logging interface.
type Logger interface {
	Errorf(string, ...any)
	Warningf(string, ...any)
	Infof(string, ...any)
	Debugf(string, ...any)
}

type SlogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger(level slog.Leveler) *SlogLogger {
	handler := slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: level,
		},
	)
	return &SlogLogger{
		logger: slog.New(handler),
	}
}

func (s *SlogLogger) Errorf(format string, args ...any) {
	s.logger.Error(format, slog.Any("args", args))
}

func (s *SlogLogger) Warningf(format string, args ...any) {
	s.logger.Warn(format, slog.Any("args", args))
}

func (s *SlogLogger) Infof(format string, args ...any) {
	s.logger.Info(format, slog.Any("args", args))
}

func (s *SlogLogger) Debugf(format string, args ...any) {
	s.logger.Debug(format, slog.Any("args", args))
}
