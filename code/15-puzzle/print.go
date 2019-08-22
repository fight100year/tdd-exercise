package main

import "fmt"

const (
	whiteBlock = "   "
	newLine    = "\n"
)

func printTable() (result string) {
	result += newLine

	for x := range table {
		for y := range table[x] {
			if table[x][y] == 0 {
				result += whiteBlock
			} else {
				result += fmt.Sprintf(" %02d", table[x][y])
			}
		}

		result += newLine
	}

	return
}
