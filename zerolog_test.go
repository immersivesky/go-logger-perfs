package logger_tests

import (
  "testing"
  "github.com/rs/zerolog"
  "github.com/rs/zerolog/log"
)

func BenchmarkZerolog(b *testing.B) {
  zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
  for i := 0; i < b.N; i++ {
    log.Info().Msg("Info text from Zerolog logger")
  }
}
