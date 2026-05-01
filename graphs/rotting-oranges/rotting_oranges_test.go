package main

import (
	"testing"
)

func TestOrangesRotting(t *testing.T) {
	// Estrutura do nosso Table-Driven Test
	testes := []struct {
		nome     string
		grid     [][]int
		esperado int
	}{
		{
			nome: "Caso 1: Todas apodrecem (Caminho Feliz)",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			esperado: 4,
		},
		{
			nome: "Caso 2: Laranja isolada (Impossível)",
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			esperado: -1,
		},
		{
			nome: "Caso 3: Já começa sem laranjas frescas",
			grid: [][]int{
				{0, 2},
			},
			esperado: 0,
		},
		{
			nome:     "Caso 4: Matriz vazia (Edge Case)",
			grid:     [][]int{},
			esperado: -1,
		},
		{
			nome: "Caso 5: Múltiplos pacientes zero (Multisource real)",
			grid: [][]int{
				{2, 1, 2},
				{1, 1, 1},
				{2, 1, 2},
			},
			esperado: 2,
		},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			// Criamos uma cópia profunda (deep copy) do grid para cada teste
			// pois nossa função modifica o grid original e isso sujaria os testes
			gridCopy := copyGrid(tt.grid)

			resultado := orangesRotting(gridCopy)

			if resultado != tt.esperado {
				t.Errorf("Resultado recebido: %d, mas esperava: %d", resultado, tt.esperado)
			}
		})
	}
}

// Função auxiliar para copiar a matriz e não vazar estado entre os testes
func copyGrid(grid [][]int) [][]int {
	if len(grid) == 0 {
		return [][]int{}
	}
	newGrid := make([][]int, len(grid))
	for i := range grid {
		newGrid[i] = make([]int, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}
