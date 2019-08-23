package palindrome

import (
	"testing"
)

func TestCheckPalindrome(t *testing.T) {
	testCases := []struct {
		testString string
		excepeted  bool
	}{
		{"", true},
		{"abcd", false},
		{"aaaa", true},
		{"aaa", true},
		{"123321", true},
		{"12321", true},
		{"1231", false},
		{"abA", true},
	}

	for _, x := range testCases {
		got := CheckPalindrome(x.testString)
		want := x.excepeted

		if got != want {
			t.Errorf("check failed: want %t got %t, %v", want, got, x.testString)
		}
	}
}
