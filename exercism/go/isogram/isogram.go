package isogram

import (
	"strings"
)

// IsIsogram determine if a word or phrase is an isogram
func IsIsogram(str string) bool {
	r := strings.NewReplacer("-", "", " ", "")
	str = r.Replace(str)
	str = strings.ToLower(str)

	chars := map[string]bool{}
	for _, x := range str {
		chars[string(x)] = true
	}

	return len(str) == len(chars)
}
