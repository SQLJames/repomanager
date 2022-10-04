package jlogr_test

import (
	"testing"

	"github.com/sqljames/repomanager/pkg/util/jlogr"
)

var (
	benchmarkLogger = jlogr.Logger.ILog
)

func BenchmarkInfoLogger(b *testing.B) {
	b.Run("Info Benchmark", Info)
}

func Info(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkLogger.Info("test", "method", "GET", "URL", "/metrics")
	}
}
