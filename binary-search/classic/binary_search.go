
package main

func Search(nums []int, target int) int{
	
	left, right := 0 , len(nums) - 1

	meio := left + (right - left) / 2

	for left <= right{
		if nums[meio] == target {
			return meio
		}
		if nums[meio] > target {
			right = meio  - 1 
		} 
		if nums[meio] < target {
			left = meio + 1
		}
		meio = left + (right - left) / 2
	}
	return -1
}