package main

import (
	"path/filepath"
	"testing"
)

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

func Benchmark_copyToClipboardUsingFinder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		absPath, _ := filepath.Abs("./bird.png")

		_ = copyToClipboardUsingFinder(absPath)
	}
}

func Benchmark_copyToClipboardInObjC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		absPath, _ := filepath.Abs("./bird.png")

		_ = copyToClipboardInObjC(absPath)
	}
}
