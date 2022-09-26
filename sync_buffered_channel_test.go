package gocp_test

import (
	"testing"
)

var bInt64Channel int64
var chBChannel = make(chan interface{}, 1)

func Benchmark_Int64_Buffered_Channel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		chBChannel <- nil
		bInt64Channel++
		<-chBChannel
	}
}
