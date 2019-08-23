package raindrops

import (
	"strconv"
)

// Convert convert a int to specical string
func Convert(input int) (expected string) {
	if input%3 == 0 {
		expected += "Pling"
	}
	if input%5 == 0 {
		expected += "Plang"
	}
	if input%7 == 0 {
		expected += "Plong"
	}

	if len(expected) == 0 {
		expected = strconv.Itoa(input)
	}

	return
}
