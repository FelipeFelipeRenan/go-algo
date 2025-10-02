package main

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "exemplo 1 - alvo presente",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			name:   "exemplo 2 - alvo ausente",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
		{
			name:   "array com um elemento - presente",
			nums:   []int{5},
			target: 5,
			want:   0,
		},
		{
			name:   "array com um elemento - ausente",
			nums:   []int{5},
			target: -5,
			want:   -1,
		},
		{
			name:   "alvo é o primeiro elemento",
			nums:   []int{2, 5, 7, 8, 11, 12},
			target: 2,
			want:   0,
		},
		{
			name:   "alvo é o último elemento",
			nums:   []int{2, 5, 7, 8, 11, 12},
			target: 12,
			want:   5,
		},
		{
			name:   "array vazio",
			nums:   []int{},
			target: 5,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Supondo que sua função se chame Search
			if got := Search(tt.nums, tt.target); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

