package main

func TwoSumOrdered(nums []int, target int) []int {
	var results []int
	if len(nums) == 0 {
		return append(results, 0, 0)
	}

	right := len(nums) - 1
	for left := 0; left < right; left++ {
		if nums[left]+nums[right] == target {
			results = append(results, left, right)
			return results
		}
		right--
	}

	return results
}
