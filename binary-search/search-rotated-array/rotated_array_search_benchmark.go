package main

import "testing"

func createRotatedSlice(size int) []int {
	slice := make([]int, size)
	pivot := size / 3
	// Preenche a parte depois do pivô
	for i := 0; i < size-pivot; i++ {
		slice[i] = pivot + i
	}
	// Preenche a parte antes do pivô
	for i := 0; i < pivot; i++ {
		slice[size-pivot+i] = i
	}
	return slice
}

var rotatedSlice = createRotatedSlice(1000000)
var targetForRotatedBenchmark = 500

func BenchmarkSearchRotated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SearchRotated(rotatedSlice, targetForRotatedBenchmark)
	}
}