package main

import "math"

func MinSubArrayLen(target int, nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	left := 0

	minLenght := math.MaxInt
	k := 0

	for right := range(nums) {
		k += nums[right]

		for k >= target {
			if right-left+1 < minLenght {
				minLenght = right - left + 1
			}
			k -= nums[left]
			left++
		}
	}

	if minLenght == math.MaxInt{
		minLenght = 0
	}
	return minLenght

}
