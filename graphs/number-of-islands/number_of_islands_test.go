package main

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]byte
		expected int
	}{
		{
			name: "Exemplo 1 - Uma grande ilha",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			name: "Exemplo 2 - Três ilhas isoladas",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
		{
			name:     "Matriz vazia",
			grid:     [][]byte{},
			expected: 0,
		},
		{
			name: "Tudo é água",
			grid: [][]byte{
				{'0', '0'},
				{'0', '0'},
			},
			expected: 0,
		},
		{
			name: "Grade xadrez (várias ilhas de tamanho 1)",
			grid: [][]byte{
				{'1', '0', '1'},
				{'0', '1', '0'},
				{'1', '0', '1'},
			},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NumberOfIslands(tt.grid)
			if got != tt.expected {
				t.Errorf("NumIslands() = %v, esperado %v", got, tt.expected)
			}
		})
	}
}
