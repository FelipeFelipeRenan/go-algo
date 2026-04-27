package main

func FindDuplicate(nums []int) int {
	slow, fast := 0, 0

	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break

		}
	}
	slow = 0
	for {
		slow = nums[slow]
		fast = nums[fast]
		if slow == fast {
			return slow
		}
	}
}

// [1,2,3,4,4,5,6]
