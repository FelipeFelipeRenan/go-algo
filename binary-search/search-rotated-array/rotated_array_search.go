package main

func SearchRotated(nums []int, target int) int {

    if len(nums) == 1 {
        if nums[0] == target {
            return 0
        } else {
            return -1
        }
    }

	l, r := 0, len(nums)-1

	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		} 
		if nums[l] <= nums[m] {
			if target >= nums[l] && target < nums[m]{
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			if target <= nums[r] && target > nums[m] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}
