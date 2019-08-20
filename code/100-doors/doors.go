package main

import (
	"fmt"
)

// Doors is 100 doors
type Doors struct {
	state []bool
	count uint
}

func (d *Doors) setDoorCount(count uint) {
	d.count = count
	d.state = make([]bool, count+1)
}

func (d *Doors) do() {
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
	d.do()

	for i, isOpened := range d.state {
		if isOpened {
			resutl = append(resutl, i)
		}
	}

	return
}

func create100Doors() *Doors {
	var doors Doors
	doors.setDoorCount(100)
	return &doors
}

func main() {
	fmt.Println(create100Doors().Pass())
}
