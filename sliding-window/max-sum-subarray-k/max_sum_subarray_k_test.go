package main

import "testing"

func TestFindMaxSumSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "exemplo 1",
			nums: []int{2, 1, 5, 1, 3, 2},
			k:    3,
			want: 9, // subarray [5, 1, 3]
		},
		{
			name: "exemplo 2 com negativos",
			nums: []int{1, 9, -1, -2, 7, 3, -1, 2},
			k:    4,
			want: 13, // subarray [9, -1, -2, 7]
		},
		{
			name: "k igual ao tamanho do array",
			nums: []int{1, 2, 3, 4, 5},
			k:    5,
			want: 15, // o próprio array
		},
		{
			name: "array com todos negativos",
			nums: []int{-1, -2, -3, -4, -5},
			k:    2,
			want: -3, // subarray [-1, -2]
		},
		{
			name: "k é 1",
			nums: []int{10, 5, 20, 8},
			k:    1,
			want: 20, // o maior elemento individual
		},
		{
			name: "k maior que o tamanho do array",
			nums: []int{1, 2, 3},
			k:    4,
			want: 0, // ou algum outro valor de erro/padrão
		},
		{
			name: "array vazio",
			nums: []int{},
			k:    2,
			want: 0, // ou algum outro valor de erro/padrão
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Supondo que sua função se chame FindMaxSumSubarray
			// Se você nomear de forma diferente, apenas ajuste aqui.
			got := FindMaxSumSubarray(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("FindMaxSumSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}