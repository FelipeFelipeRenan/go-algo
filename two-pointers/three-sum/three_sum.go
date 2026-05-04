package main

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right]
			if sum+nums[i] == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && right < len(nums) && nums[right] == nums[right+1] {

					right--
				}
				continue
			}
			if sum < -nums[i] {
				left++
			}
			if sum > -nums[i] {
				right--
			}
		}
	}
	return result
}
