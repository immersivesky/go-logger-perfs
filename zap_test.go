package logger_test

import (
	"testing"

	"go.uber.org/zap"
)

func BenchmarkZap(b *testing.B) {
	logger, _ := zap.NewProduction()

	for i := 0; i < b.N; i++ {
		logger.Info("Info text from Zap logger")
	}
}
