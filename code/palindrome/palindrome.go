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
