package main

import (
	"testing"
)

var fa int

func BenchmarkBST(b testing.B) {
	var a int

	for i := 0; i < b.N; i++ {
		a = main()
	}

	fa = a
}
