package main


func removeElement(nums []int, val int) int {

	if len(nums) == 0 {
		return 0
	}

	left := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[left] = nums[i]
			left++

		}
	}
	return left
}