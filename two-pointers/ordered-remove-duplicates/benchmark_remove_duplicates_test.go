package main

import "testing"


func BenchmarkRemoveDuplicates(b *testing.B) {
	// Um array grande com muitas duplicatas
	nums := make([]int, 0, 100000)
	for i := 0; i < 10000; i++ {
		for j := 0; j < 10; j++ {
			nums = append(nums, i)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Criar uma cópia do array para cada iteração
		cpy := make([]int, len(nums))
		copy(cpy, nums)
		RemoveDuplicates(cpy)
	}
}
