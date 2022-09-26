package main

import "testing"

var bIntNoSync int64

func Benchmark_Int64_NoSync(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bIntNoSync++
	}
}
