package logger_test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func BenchmarkLogrus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logrus.Info("Info text from Logrus logger")
	}
}
