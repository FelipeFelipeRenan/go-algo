package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		input          []int
		expectedK      int
		expectedUnique []int
	}{
		{
			name:           "array vazio",
			input:          []int{},
			expectedK:      0,
			expectedUnique: []int{},
		},
		{
			name:           "sem duplicatas",
			input:          []int{1, 2, 3, 4},
			expectedK:      4,
			expectedUnique: []int{1, 2, 3, 4},
		},
		{
			name:           "com duplicatas simples",
			input:          []int{1, 1, 2},
			expectedK:      2,
			expectedUnique: []int{1, 2},
		},
		{
			name:           "com duplicatas múltiplas",
			input:          []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedK:      5,
			expectedUnique: []int{0, 1, 2, 3, 4},
		},
		{
			name:           "todos iguais",
			input:          []int{2, 2, 2, 2},
			expectedK:      1,
			expectedUnique: []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			k := RemoveDuplicates(inputCopy)

			if k != tt.expectedK {
				t.Errorf("esperado k=%d, obtido k=%d", tt.expectedK, k)
			}

			if !reflect.DeepEqual(inputCopy[:k], tt.expectedUnique) {
				t.Errorf("esperado slice %v, obtido %v", tt.expectedUnique, inputCopy[:k])
			}
		})
	}
}
