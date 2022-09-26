package main

import (
	"sync"
	"testing"
)

var bInt64Mutex int64
var mInt64Mutex sync.Mutex

func Benchmark_Int64_Mutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mInt64Mutex.Lock()
		bInt64Mutex++
		mInt64Mutex.Unlock()
	}
}
