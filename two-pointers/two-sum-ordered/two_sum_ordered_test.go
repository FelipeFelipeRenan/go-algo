package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSumOrdered(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"valores positivos - válido", []int{1, 2, 3, 5, 9}, 8, []int{2, 3}},    // 3 + 5 = 8
		{"com negativos - válido", []int{-10, -3, -1, -20}, -4, []int{2, 3}},   // -3 + -1 = -4
		{"único elemento", []int{42}, 42, nil},                                 // impossível
		{"vazio", []int{}, 0, []int{0, 0}},                                     // padrão definido
		{"valores iguais - válido", []int{7, 7, 7}, 14, []int{0, 2}},           // 7 + 7 = 14
		{"sem solução", []int{1, 2, 3}, 100, nil},                              // impossível
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorted := make([]int, len(tt.nums))
			copy(sorted, tt.nums)
			sort.Ints(sorted)

			got := TwoSumOrdered(sorted, tt.target)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FAIL %s: got %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
