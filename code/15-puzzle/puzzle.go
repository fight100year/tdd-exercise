package main

import (
	"fmt"
	"os"
)

var (
	countStep int
)

func show() {
	fmt.Printf("total step: %d\n", countStep)
	countStep++
	fmt.Println(printTable())
}

func run() {
	startCaptureKeyBoard()
	randomTable()
	for {
		select {
		case key := <-steps:
			switch key {
			case ESC:
				os.Exit(0)
			case UP:
				fallthrough
			case DOWN:
				fallthrough
			case LEFT:
				fallthrough
			case RIGHT:
				move(key)
				break
			default:
			}
		}
	}
}

func main() {
	run()
}
