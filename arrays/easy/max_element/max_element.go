package main


func MaxElement(nums []int) int  {

	if len(nums) == 0{
		return 0
	}
	max := nums[0] 
	for _, v := range nums {
		if v > max{
			max = v
		}
	}
	return max
}