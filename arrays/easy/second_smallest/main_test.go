package main

import "testing"

func TestSecondSmallest(t *testing.T) {
    tests := []struct {
        name string
        nums []int
        want int
    }{
        {"valores distintos", []int{3, 1, 2, 2, 1, 5}, 2},
        {"apenas repetidos", []int{5, 5, 5, 5}, 5},
        {"negativos com repetidos", []int{-10, -5, -5, -1}, -5},
        {"um único valor", []int{1}, 1},
        {"crescente", []int{1, 2, 3, 4, 5}, 2},
        {"decrescente", []int{5, 4, 3, 2, 1}, 2},
        {"com negativos iguais", []int{-1, -1, -1, -1}, -1},
        {"menor valor duplicado com próximo", []int{-2, -2, -1}, -1},
        {"menor distinto no final", []int{10, 10, 10, 5}, 10},
        {"vazio", []int{}, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := SecondSmallest(tt.nums)
            if got != tt.want {
                t.Errorf("got %d, want %d", got, tt.want)
            }
        })
    }
}
