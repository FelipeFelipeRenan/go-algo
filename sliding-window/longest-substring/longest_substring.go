package main

func LengthOfLongestSubstring(s string) int {

	if len(s) <= 1 {
		return len(s)
	}

	charIndexMap := make(map[rune]int)

	left := 0
	maxLenght := 0

	for right, currentChar := range s {

		if lastSeenIndex, ok := charIndexMap[currentChar]; ok && lastSeenIndex+1 > left {

			left = lastSeenIndex + 1
		}

		charIndexMap[currentChar] = right

		currentLenght := right - left + 1

		if currentLenght > maxLenght {
			maxLenght = currentLenght
		}

		right++
	}

	return maxLenght
}
