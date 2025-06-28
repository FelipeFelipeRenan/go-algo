package main

import "testing"

func TestSumOfEvens(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want int
    }{
        {"exemplo básico", []int{1, 2, 3, 4, 5, 6}, 12},
        {"nenhum par", []int{1, 3, 5}, 0},
        {"todos pares", []int{2, 4, 6}, 12},
        {"vazio", []int{}, 0},
        {"com negativos", []int{-2, -4, 1}, -6},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := SumOfEvens(tt.nums)
            if got != tt.want {
                t.Errorf("got %d, want %d", got, tt.want)
            }
        })
    }
}
