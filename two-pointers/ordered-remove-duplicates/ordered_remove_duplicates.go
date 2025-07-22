package main

func RemoveDuplicates(numbers []int) int {

	if len(numbers) <= 1 {
		return len(numbers)
	}

	//var size int = 1
	left, right, size := 0, 1, 1

	for right < len(numbers) {

		if numbers[left] != numbers[right] {
			left++
			numbers[left] = numbers[right]
			size++
		} else {
			right++
		}
	}

	return size
}
