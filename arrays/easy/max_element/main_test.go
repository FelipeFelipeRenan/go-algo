package main

import "testing"

func TestMaxElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"valores positivos", []int{3, 5, 1, 9, 2}, 9},
		{"com negativos", []int{-10, -3, -1, -20}, -1},
		{"único elemento", []int{42}, 42},
		{"vazio", []int{}, 0}, // aqui decidimos retornar 0 como padrão
		{"valores iguais", []int{7, 7, 7}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MaxElement(tt.nums)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}
