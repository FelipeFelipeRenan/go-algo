package main

func CutRod(prices []int, size int) int {

	memo := make([]int, size+1)
	for i := range memo {
		memo[i] = -1
	}
	return cutRodHelper(prices, size, memo)
}

func cutRodHelper(prices []int, size int, memo []int) int {
	if size == 0 {
		return 0
	}

	if memo[size] != -1 {
		return memo[size]
	}

	dp := -1

	for i := 0; i < len(prices) && i+1 <= size; i++ {
		tamanhoCorte := i + 1
		tamanhoSobre := size - tamanhoCorte
		lucroAtual := prices[i] + cutRodHelper(prices, tamanhoSobre, memo)

		if lucroAtual > dp {
			dp = lucroAtual
		}
	}

	memo[size] = dp
	return dp
}
