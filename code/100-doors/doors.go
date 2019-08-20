package main

import (
	"fmt"
)

// Door output the final state of 100 doors
func Door() []int {
	return configDoor(100)
}

func changeDoor(state int) int {
	if state == 1 {
		return 0
	}

	return 1
}

func configDoor(countDoor int) []int {
	doors := make([]int, countDoor+1)
	for i := 1; i <= countDoor; i++ {
		j := i
		for j <= countDoor {
			doors[j] = changeDoor(doors[j])
			j += i
		}
	}

	return doors[1:]
}

func main() {
	fmt.Printf("%v", Door())
}
