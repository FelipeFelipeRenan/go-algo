package main

// import "math"

func FindMaxSumSubarray(values []int, k int) int {

	if len(values) == 0 || k > len(values) {
		return 0
	}
	if len(values) == 1 {
		return values[0]
	}
	if len(values) == k {
		sum := 0
		for _, v := range values {
			sum = sum + v
		}
		return sum
	}

	sum := 0

	for i := 0; i < k; i++ {
		sum = sum + values[i]
	}	

	currentSum := sum

	left := 1
	right := k
	for right < len(values) {

		currentSum = currentSum - values[left - 1] + values[right]

		if currentSum > sum {
			sum = currentSum
		}
		left++
		right++
	}

	return sum
}
