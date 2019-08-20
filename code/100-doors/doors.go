package main

import (
	"fmt"
)

// Door output the final state of 100 doors
func Door() []int {
	return configDoor(100)
}

func configDoor(countDoor int) []int {
	return []int{}
}

func main() {
	fmt.Printf("%v", Door())
}
