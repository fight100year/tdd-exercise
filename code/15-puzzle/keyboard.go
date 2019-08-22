package main

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type enterKey int

// key
const (
	UP enterKey = iota
	DOWN
	LEFT
	RIGHT
	ESC
)

var (
	steps chan enterKey
)

func init() {
	steps = make(chan enterKey)
}

func startCaptureKeyBoard() {
	go func() {
		if err := termbox.Init(); err != nil {
			panic(err)
		}
		defer termbox.Close()

		for i := 0; i < 80; i++ {
			fmt.Println()
		}
		show()

		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				{
					switch ev.Key {
					case termbox.KeyEsc:
						steps <- ESC
						break
					case termbox.KeyArrowUp:
						steps <- DOWN
						break
					case termbox.KeyArrowDown:
						steps <- UP
						break
					case termbox.KeyArrowLeft:
						steps <- RIGHT
						break
					case termbox.KeyArrowRight:
						steps <- LEFT
						break
					}
				}
			}
		}
	}()
}
