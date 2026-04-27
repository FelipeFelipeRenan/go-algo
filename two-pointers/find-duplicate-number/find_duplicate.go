package main

func FindDuplicate(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	slow := 0
	fast := 1

	for slow < len(nums) {
		for fast < len(nums) {
			if nums[slow] != nums[fast] {
				fast++
			} else {
				return nums[slow]
			}
		}
		slow++
		fast = slow + 1
	}
	return -1
}

// [1,2,3,4,4,5,6]
