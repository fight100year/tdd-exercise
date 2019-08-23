package palindrome

import (
	"strings"
)

// CheckPalindrome check a is palindrome or not
func CheckPalindrome(testString string) bool {
	testString = strings.ToLower(testString)

	reverseString := ""
	for _, x := range testString {
		reverseString = string(x) + reverseString
	}

	return reverseString == testString
}

// CheckPalindrome2 is another impl
func CheckPalindrome2(testString string) bool {
	testSlice := []rune(strings.ToLower(testString))

	lenght := len(testSlice)

	for i := 0; i <= lenght/2; i++ {
		if testSlice[i] != testSlice[lenght-1-i] {
			return false
		}
	}

	return true
}
