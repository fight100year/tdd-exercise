package main

import (
	"fmt"
)

// Doors is 100 doors
type Doors struct {
	state []bool
	count uint
}

// SetCount set the number of doors
func (d *Doors) SetCount(count uint) {
	d.count = count
	d.state = make([]bool, count+1)
}

func (d *Doors) changeState() {
	for i := uint(1); i <= d.count; i++ {
		j := i
		for j <= d.count {
			d.state[j] = !d.state[j]
			j += i
		}
	}
}

// Pass query the opened door
func (d *Doors) Pass() (resutl []int) {
	d.changeState()

	for i, isOpened := range d.state {
		if isOpened {
			resutl = append(resutl, i)
		}
	}

	return
}

func query100DoorState() []int {
	var doors Doors
	doors.SetCount(100)
	return doors.Pass()
}

func main() {
	fmt.Println(query100DoorState())
}
