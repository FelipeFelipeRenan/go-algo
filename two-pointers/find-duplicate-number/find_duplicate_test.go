package main

import "testing"

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "exemplo 1",
			nums: []int{1, 3, 4, 2, 2},
			want: 2,
		},
		{
			name: "exemplo 2 - duplicata no inicio",
			nums: []int{3, 1, 3, 4, 2},
			want: 3,
		},
		{
			name: "todos os elementos iguais",
			nums: []int{2, 2, 2, 2, 2},
			want: 2, // O ciclo é imediato
		},
		{
			name: "array minimo possivel",
			nums: []int{1, 1},
			want: 1, // nums tem tamanho n+1 (2), elementos no intervalo [1, 1]
		},
		{
			name: "duplicata adjacente no final",
			nums: []int{1, 2, 3, 4, 4},
			want: 4,
		},
		{
			name: "ciclo longo no array",
			nums: []int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplicate(tt.nums); got != tt.want {
				t.Errorf("FindDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
