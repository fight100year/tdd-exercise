package hamming

import (
	"errors"
)

// Distance calc the diff of two strings
func Distance(a, b string) (int, error) {
	if len(a) == 0 && len(b) == 0 {
		return 0, nil
	}

	if len(a) != len(b) {
		return 0, errors.New("lenght is diffent")
	}

	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}

	return diff, nil
}
