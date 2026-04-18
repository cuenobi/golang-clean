package logger

import (
	"os"
	"strings"

	"github.com/cuenobi/golang-clean/internal/shared/config"
	"github.com/rs/zerolog"
)

type Logger interface {
	Debug(msg string, fields map[string]any)
	Info(msg string, fields map[string]any)
	Warn(msg string, fields map[string]any)
	Error(msg string, err error, fields map[string]any)
}

type structuredLogger struct {
	log zerolog.Logger
}

func New(cfg config.Config) Logger {
	level := zerolog.InfoLevel
	if parsed, err := zerolog.ParseLevel(strings.ToLower(cfg.LogLevel)); err == nil {
		level = parsed
	}

	base := zerolog.New(os.Stdout).
		Level(level).
		With().
		Timestamp().
		Str("service", cfg.AppName).
		Str("env", cfg.AppEnv).
		Logger()

	return &structuredLogger{log: base}
}

func (l *structuredLogger) Info(msg string, fields map[string]any) {
	ev := l.log.Info()
	for key, val := range fields {
		ev = ev.Interface(key, val)
	}
	ev.Msg(msg)
}

func (l *structuredLogger) Debug(msg string, fields map[string]any) {
	ev := l.log.Debug()
	for key, val := range fields {
		ev = ev.Interface(key, val)
	}
	ev.Msg(msg)
}

func (l *structuredLogger) Warn(msg string, fields map[string]any) {
	ev := l.log.Warn()
	for key, val := range fields {
		ev = ev.Interface(key, val)
	}
	ev.Msg(msg)
}

func (l *structuredLogger) Error(msg string, err error, fields map[string]any) {
	ev := l.log.Error()
	if err != nil {
		ev = ev.Err(err)
	}
	for key, val := range fields {
		ev = ev.Interface(key, val)
	}
	ev.Msg(msg)
}
