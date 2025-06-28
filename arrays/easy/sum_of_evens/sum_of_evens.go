package main

func SumOfEvens(num []int) int {

	acc := 0
	for _, v := range num {
		if v%2 == 0 {
			acc = acc + v
		}
	}
	return acc
}
