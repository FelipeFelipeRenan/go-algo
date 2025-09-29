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
			size = 0
			left++
		}
		size++
		right++
	}


	return size

}
