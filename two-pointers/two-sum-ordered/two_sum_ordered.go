package main

func TwoSumOrdered(nums []int, target int) []int {
	var results []int
	if len(nums) == 0 {
		return append(results, 0, 0)
	}

	left := 0
	right := len(nums) - 1
	
	for left < right{
		sum := nums[left]+nums[right] 
		if sum > target {
			right--
		}
		if sum < target {
			left++
		}
		if sum == target {
			results = append(results, left, right)
			return results
		}

	}

	return results
}
