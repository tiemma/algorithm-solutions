package main

import "testing"

func BenchmarkMyFunction(b *testing.B) {
	run(1_000_000, 60)
}
