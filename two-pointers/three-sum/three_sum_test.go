package main

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	testes := []struct {
		nome     string
		entrada  []int
		esperado [][]int
	}{
		{
			nome:     "Caso 1: Exemplo clássico do LeetCode",
			entrada:  []int{-1, 0, 1, 2, -1, -4},
			esperado: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nome:     "Caso 2: Sem nenhuma combinação possível",
			entrada:  []int{0, 1, 1},
			esperado: [][]int{}, // Não soma zero de jeito nenhum
		},
		{
			nome:     "Caso 3: Todos os números iguais a zero",
			entrada:  []int{0, 0, 0},
			esperado: [][]int{{0, 0, 0}},
		},
		{
			nome:     "Caso 4: Lidando com muitas duplicatas (A grande pegadinha)",
			entrada:  []int{-2, 0, 0, 2, 2},
			esperado: [][]int{{-2, 0, 2}},
		},
		{
			nome:     "Caso 5: Array gigante de zeros",
			entrada:  []int{0, 0, 0, 0, 0},
			esperado: [][]int{{0, 0, 0}}, // Deve retornar apenas um trio de zeros
		},
	}

	for _, tt := range testes {
		t.Run(tt.nome, func(t *testing.T) {
			// Criamos uma cópia do array de entrada, pois o algoritmo original
			// geralmente faz nums.Sort(), o que modificaria a variável do teste original
			entradaCopia := make([]int, len(tt.entrada))
			copy(entradaCopia, tt.entrada)

			resultado := threeSum(entradaCopia)

			// Tratamento para evitar que o Go falhe ao comparar nil com matriz vazia []
			if len(resultado) == 0 && len(tt.esperado) == 0 {
				return // Passou!
			}

			if !reflect.DeepEqual(resultado, tt.esperado) {
				t.Errorf("\nRecebido: %v\nEsperado: %v", resultado, tt.esperado)
			}
		})
	}
}
