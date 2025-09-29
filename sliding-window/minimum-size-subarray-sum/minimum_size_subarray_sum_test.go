package main

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{
			name:   "exemplo 1",
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			want:   2, // [4, 3]
		},
		{
			name:   "exemplo 2",
			target: 4,
			nums:   []int{1, 4, 4},
			want:   1, // [4]
		},
		{
			name:   "exemplo 3 - sem solução",
			target: 11,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			want:   0,
		},
		{
			name:   "soma exata no final",
			target: 8,
			nums:   []int{3, 4, 1, 1, 6},
			want:   3, // [1, 1, 6]
		},
		{
			name:   "array inteiro é a única solução",
			target: 15,
			nums:   []int{1, 2, 3, 4, 5},
			want:   5,
		},
		{
			name:   "array vazio",
			target: 100,
			nums:   []int{},
			want:   0,
		},
		{
			name:   "único elemento que satisfaz",
			target: 5,
			nums:   []int{10},
			want:   1,
		},
		{
			name:   "único elemento que não satisfaz",
			target: 10,
			nums:   []int{5},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Supondo que sua função se chame MinSubArrayLen
			// Se nomear de forma diferente, apenas ajuste aqui.
			if got := MinSubArrayLen(tt.target, tt.nums); got != tt.want {
				t.Errorf("MinSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}