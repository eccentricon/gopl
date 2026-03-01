package main

import (
	"testing"
)

func BenchmarkEcho(b *testing.B) {
	// Setup 10000 arguments
	scale := 10_000
	args := make([]string, scale)
	for i := 0; i < scale; i++ {
		args[i] = "argument"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Echo(args)
	}
}
