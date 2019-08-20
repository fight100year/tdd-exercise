package main

import (
	"fmt"
)

func init() {

}

var (
	table = [16]int{}
)

func createTable() string {
	return `
 01
`
}

func printTable() (result string) {
	result += fmt.Sprintln()
	for i, element := range table {
		if element == 0 {
			result += "   "
		} else {
			result += fmt.Sprintf(" %02d", element)
		}

		if (i+1)%4 == 0 {
			result += fmt.Sprintln()
		}
	}

	return
}

func main() {

}
