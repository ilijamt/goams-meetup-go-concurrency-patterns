package gocp_test

import (
	"sync/atomic"
	"testing"
)

var bInt64Atomic int64

func BenchmarkAtomic_AddInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&bInt64Atomic, 1)
	}
}
