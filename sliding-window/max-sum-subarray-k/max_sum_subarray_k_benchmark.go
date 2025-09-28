package main

import (
	"math/rand"
	"testing"
)

// generateRandomSlice cria um slice grande com números aleatórios para o benchmark.
func generateRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(100) - 50 // Números entre -50 e 49
	}
	return slice
}

var largeSlice = generateRandomSlice(100000)
var kBenchmark = 1000

func BenchmarkFindMaxSumSubarray_SlidingWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Altere o nome da função para corresponder à sua implementação da sliding window
		FindMaxSumSubarray(largeSlice, kBenchmark)
	}
}

// Você pode adicionar um segundo benchmark para comparar com a força bruta
/*
func BenchmarkFindMaxSumSubarray_BruteForce(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // Altere o nome da função para corresponder à sua implementação de força bruta
        FindMaxSumSubarrayBruteForce(largeSlice, kBenchmark)
    }
}
*/
