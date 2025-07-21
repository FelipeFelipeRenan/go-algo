package main

import "math"

func SecondSmallest(nums []int) int{
	
	if len(nums) == 0{
		return 0
	}
	if len(nums) == 1{
		return nums[0]
	}

	menor := math.MaxInt64
	segundo_menor := math.MaxInt64

	for _, v := range nums {
		if v < menor{
			segundo_menor = menor
			menor = v
		} else if v > menor && v < segundo_menor{
			segundo_menor = v
		} 
	}

	if segundo_menor == math.MaxInt64{
		return menor
	}

	return segundo_menor
}
