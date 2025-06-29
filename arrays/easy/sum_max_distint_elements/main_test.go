package main

import "testing"

func TestSumTwoLargest(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want int
    }{
        {"valores distintos", []int{3, 7, 1, 9, 5}, 16},      // 9 + 7
        {"com duplicatas", []int{3, 7, 1, 9, 9, 5}, 16},      // 9 + 7
        {"todos iguais", []int{5, 5, 5, 5}, 5},               // só um valor distinto
        {"um único valor", []int{42}, 42},                   // apenas um número
        {"nenhum valor", []int{}, 0},                        // vazio
        {"valores negativos", []int{-10, -5, -1}, -6},        // -1 + -5
        {"repetições e negativos", []int{-1, -1, -2, -2}, -3}, // -1 + -2
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := SumTwoLargest(tt.nums)
            if got != tt.want {
                t.Errorf("got %d, want %d", got, tt.want)
            }
        })
    }
}
