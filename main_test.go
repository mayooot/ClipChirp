package main

import "testing"

func Benchmark_resolveAbsolutePathUsingRealpath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = resolveAbsolutePathUsingRealpath("./bird.png")
	}

}

func Benchmark_resolveAbsolutePathInGo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = resolveAbsolutePathInGo("./bird.png")
	}
}
