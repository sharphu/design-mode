package chain

import (
	"testing"
)

// TestLoggerChain 测试 Logger 责任链的行为
func TestLoggerChain(t *testing.T) {
	// 创建日志责任链
	infoLogger := NewInfoLogger()
	debugLogger := NewDebugLogger()
	errorLogger := NewErrorLogger()

	infoLogger.SetNext(debugLogger)
	debugLogger.SetNext(errorLogger)

	// 测试 InfoLevel 日志，确保责任链调用到 InfoLogger
	t.Run("Info Level Logging", func(t *testing.T) {
		infoLogger.Log(InfoLevel, "Testing info level log")
	})

	// 测试 DebugLevel 日志，确保责任链调用到 DebugLogger
	t.Run("Debug Level Logging", func(t *testing.T) {
		infoLogger.Log(DebugLevel, "Testing debug level log")
	})

	// 测试 ErrorLevel 日志，确保责任链调用到 ErrorLogger
	t.Run("Error Level Logging", func(t *testing.T) {
		infoLogger.Log(ErrorLevel, "Testing error level log")
	})
}
