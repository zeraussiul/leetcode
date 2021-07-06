package palindrome

import (
	"strconv"
)

// algo using strconv
func isPalindrome(x int) bool {
	// if negative, not palindrome
	if x < 0 {
		return false
	}
	// convert to string then to []rune to iterate thru string easier
	//num := []rune(strconv.Itoa(x))
	num := strconv.Itoa(x)

	for i, j := 0, len(num)-1; i <= j; i, j = i+1, j-1 {
		if num[i] != num[j] {
			return false
		}
	}
	return true
}

// do it with math instead
