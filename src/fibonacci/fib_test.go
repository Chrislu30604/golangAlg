package main

import "testing"

func BenchmarkFib(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib(30)
	}
}

func BenchmarkFib1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib2(30)
	}
}
func BenchmarkFib2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib3(30)
	}
}
