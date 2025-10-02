package main

import (
	"testing"
)

func createSortedSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i * 2 // Apenas um exemplo de preenchimento ordenado
	}
	return slice
}

var sortedSlice = createSortedSlice(1000000)
var targetForClassicBenchmark = 987654

func BenchmarkSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Search(sortedSlice, targetForClassicBenchmark)
	}
}

// Para comparação, um benchmark de busca linear
func linearSearch(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
	}
	return -1
}

func BenchmarkLinearSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		linearSearch(sortedSlice, targetForClassicBenchmark)
	}
}