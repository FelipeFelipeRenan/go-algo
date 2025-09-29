package main

func LengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	charIndexMap := make(map[rune]int)

	left, right := 0, 0
	maxLenght := 0
	
	for right < len(s) {
		
		currentChar := rune(s[right])
		
		lastSeenIndex , ok := charIndexMap[currentChar]
		
		if ok{
			if lastSeenIndex + 1 > left{
				
				left = lastSeenIndex + 1
			}
		}
		
		currentLenght := right - left + 1
		charIndexMap[currentChar] = right
		
		if currentLenght > maxLenght{
			maxLenght = currentLenght
		}

		right++
	}

	return maxLenght
}

// abcabb