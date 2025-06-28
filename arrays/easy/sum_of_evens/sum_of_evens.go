package main

func SumOfEvens(num []int) int {

	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	acc := 0
	for _, v := range numeros {
		if v%2 == 0 {
			acc = acc + v
		}
	}
	return (acc)
}
