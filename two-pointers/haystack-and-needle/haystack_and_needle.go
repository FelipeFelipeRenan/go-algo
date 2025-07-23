package main

func strStr(haystack string, needle string) int {
    
    if len(needle) == 0{
        return -1
    } 

    left := 0
    for i := 0; i < len(haystack); i++{

        if haystack[i] == needle[left] {
                left++
            if left == len(needle) {
                return i - left + 1
            }
        }  else {
            i = i - left
            left = 0
        }
        
    }
    return -1
}