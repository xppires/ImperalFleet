package interfaces

import ( 
	"log/slog"
	"os"
)

type Logger interface {
    Info(msg string, fields ...interface{})
    Warn(msg string, fields ...interface{})
    Error(msg string, fields ...interface{})
    Debug(msg string, fields ...interface{})
}

// SlogLogger implements Logger using slog
type SlogLogger struct {
    logger *slog.Logger
}

// NewSlogLogger creates a new SlogLogger with JSON output
func NewSlogLogger() *SlogLogger {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level:     slog.LevelInfo,
        AddSource: true,
    }))
	slog.SetDefault(logger)
    return &SlogLogger{logger: logger}
}

// Info implements Logger.Info
func (l *SlogLogger) Info(msg string, fields ...interface{}) {
    l.logger.Info(msg, fields...)
}

// Warn implements Logger.Warn
func (l *SlogLogger) Warn(msg string, fields ...interface{}) {
    l.logger.Warn(msg, fields...)
}

// Error implements Logger.Error
func (l *SlogLogger) Error(msg string, fields ...interface{}) {
    l.logger.Error(msg, fields...)
}

// Debug implements Logger.Debug
func (l *SlogLogger) Debug(msg string, fields ...interface{}) {
    l.logger.Debug(msg, fields...)
}