package main

import "testing"

var fa int

func BenchmarkSum(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = sum(i)
	}

	fa = a
}

func BenchmarkQuickSum(b *testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = quickSum(i)
	}

	fa = a
}
