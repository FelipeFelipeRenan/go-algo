package main

import "testing"

func TestCutRod(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		n        int
		expected int
	}{
		{
			name:     "Exemplo 1 Clássico (Tamanho 2 e 6)",
			prices:   []int{1, 5, 8, 9, 10, 17, 17, 20},
			n:        8,
			expected: 22,
		},
		{
			name:     "Exemplo 2 (Cortar tudo em tamanho 1)",
			prices:   []int{3, 5, 8, 9, 10, 17, 17, 20},
			n:        8,
			expected: 24,
		},
		{
			name:     "Haste de tamanho 4",
			prices:   []int{1, 5, 8, 9, 10, 17, 17, 20},
			n:        4,
			expected: 10, // Cortar em dois de tamanho 2 (5 + 5 = 10)
		},
		{
			name:     "Não cortar é o mais lucrativo",
			prices:   []int{1, 1, 1, 10},
			n:        4,
			expected: 10,
		},
		{
			name:     "Haste de tamanho 0",
			prices:   []int{1, 5, 8},
			n:        0,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CutRod(tt.prices, tt.n)
			if got != tt.expected {
				t.Errorf("CutRod() = %v, esperado %v", got, tt.expected)
			}
		})
	}
}
