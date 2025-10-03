package main

import "testing"

func TestSearchRotated(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "exemplo 1",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "exemplo 2 - alvo ausente",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			name:   "exemplo 3 - array pequeno",
			nums:   []int{1},
			target: 0,
			want:   -1,
		},
		{
			name:   "sem rotação - alvo presente",
			nums:   []int{1, 2, 3, 4, 5, 6},
			target: 4,
			want:   3,
		},
		{
			name:   "alvo na parte rotacionada",
			nums:   []int{5, 1, 3},
			target: 3,
			want:   2,
		},
		{
			name:   "alvo é o primeiro elemento",
			nums:   []int{3, 4, 5, 1, 2},
			target: 3,
			want:   0,
		},
		{
			name:   "alvo é o último elemento",
			nums:   []int{7, 8, 1, 2, 3, 4, 5, 6},
			target: 6,
			want:   7,
		},
		{
			name:   "array com dois elementos - rotação",
			nums:   []int{3, 1},
			target: 1,
			want:   1,
		},
		{
			name:   "array com dois elementos - sem rotação",
			nums:   []int{1, 3},
			target: 1,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Nome da função sugerido: SearchRotated
			if got := SearchRotated(tt.nums, tt.target); got != tt.want {
				t.Errorf("SearchRotated() = %v, want %v", got, tt.want)
			}
		})
	}
}