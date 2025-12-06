package main

func TwoSum(nums []int, target int) []int{

	seen := make(map[int]int)

	for i, v := range nums {
		complement := target - v

		if indexOfComplement , ok := seen[complement]; ok {
			return []int{indexOfComplement, i}
		}
		seen[v] = i
	}

	return nil
}