package main

func MinSubArrayLen(target int, nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	size := 0

	left, right := 0, 0

	k := 0

	for right < len(nums) {
		k += nums[right]

		if k > target{
			left++
		}
		if target != k {

			right++
		}

		right++
	}

	return size

}
