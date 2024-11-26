package logger

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zapLogger *zap.Logger
	config    Config
	mu        sync.RWMutex
}

type Config struct {
	Level       string // Log level: debug, info, warn, error
	Environment string // Environment: production, development
}

// DefaultConfig provides a baseline configuration for the logger.
func DefaultConfig() Config {
	return Config{
		Level:       "info",
		Environment: "production",
	}
}

var (
	globalLogger *Logger
	once         sync.Once
)

// GetLogger initializes or returns the singleton logger instance.
func GetLogger() *Logger {
	once.Do(func() {
		globalLogger = &Logger{}
		globalLogger.configure(DefaultConfig())
	})
	return globalLogger
}

// configure initializes the logger based on the given config.
func (l *Logger) configure(cfg Config) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var level zapcore.Level
	switch cfg.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	var zapCfg zap.Config
	if cfg.Environment == "development" {
		zapCfg = zap.NewDevelopmentConfig()
	} else {
		zapCfg = zap.NewProductionConfig()
	}

	zapCfg.Level = zap.NewAtomicLevelAt(level)
	zapCfg.EncoderConfig.TimeKey = "timestamp"
	zapCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := zapCfg.Build()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	l.zapLogger = logger
	l.config = cfg
}

// UpdateConfig allows dynamic reconfiguration of the logger.
func (l *Logger) UpdateConfig(cfg Config) {
	l.configure(cfg)
}

// Debug logs a debug message.
func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.zapLogger.Debug(msg, fields...)
}

// Info logs an informational message.
func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.zapLogger.Info(msg, fields...)
}

// Warn logs a warning message.
func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.zapLogger.Warn(msg, fields...)
}

// Error logs an error message.
func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.zapLogger.Error(msg, fields...)
}

// Fatal logs a fatal message and exits the application.
func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.zapLogger.Fatal(msg, fields...)
}

// Sync flushes any buffered log entries.
func (l *Logger) Sync() {
	_ = l.zapLogger.Sync()
}

// LogHTTPRequest logs details of an HTTP request.
func (l *Logger) LogHTTPRequest(method, path, traceID string, statusCode int, durationMs float64) {
	l.Info("HTTP Request",
		zap.String("method", method),
		zap.String("path", path),
		zap.String("trace_id", traceID),
		zap.Int("status_code", statusCode),
		zap.Float64("duration_ms", durationMs),
	)
}

// LogTrace logs details of a trace span.
func (l *Logger) LogTrace(spanName, traceID string, durationMs float64) {
	l.Debug("Trace Span",
		zap.String("span_name", spanName),
		zap.String("trace_id", traceID),
		zap.Float64("duration_ms", durationMs),
	)
}
