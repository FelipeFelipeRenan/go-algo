package main


func ContainerWithMostWater(height []int) int{

	if len(height) == 0{
		return 0
	}
	
	left, right, area := 0, len(height) - 1, 0

	maxArea := 0
	for left < right {
		area = (right - left) * min(height[left], height[right])

		if area > maxArea{
			maxArea = area
		}
		if height[left] < height[right]{
			left++
		} else {
			right--
		}
	}

	return maxArea
}