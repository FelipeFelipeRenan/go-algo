package main

import (
	//"fmt"
	"math"
)

func SumTwoLargest(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	maxN := math.MinInt64
	maxN2 := math.MinInt64

	for _, v := range nums {
		if v > maxN {
			maxN2 = maxN
			maxN = v
		} else if v != maxN && v > maxN2 {
			maxN2 = v
		}
	}


	if maxN2 == math.MinInt64 {
		return maxN
	}

	return maxN + maxN2

}
