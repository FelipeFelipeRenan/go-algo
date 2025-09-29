package main

import (
	"math/rand"
	"testing"
	"time"
)

func generateRandomPositiveSlice(size int) []int {
	slice := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100) + 1 // Números positivos entre 1 e 100
	}
	return slice
}

var largeSliceForSum = generateRandomPositiveSlice(100000)
var targetForBenchmark = 10000

func BenchmarkMinSubArrayLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MinSubArrayLen(targetForBenchmark, largeSliceForSum)
	}
}