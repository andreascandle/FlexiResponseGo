package logger_test

import (
	"testing"

	"github.com/andreascandle/FlexiResponseGo/logger"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestLoggerInitialization(t *testing.T) {
	log := logger.GetLogger()
	assert.NotNil(t, log)

	log.Info("Test log", zap.String("key", "value"))
}

func TestLoggerDynamicConfigUpdate(t *testing.T) {
	log := logger.GetLogger()
	log.UpdateConfig(logger.Config{
		Level:       "debug",
		Environment: "development",
	})
	log.Debug("This is a debug message")
	assert.NotNil(t, log)
}
