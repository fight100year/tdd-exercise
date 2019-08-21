package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

var (
	table = [16]int{}
	steps chan enterKey
)

func init() {
	steps = make(chan enterKey)
}

const (
	whiteBlock = "   "
	newLine    = "\n"
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

func initTable() {
	for i := range table {
		table[i] = i
	}
}

func randomTable() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := len(table) - 1; i > 0; i-- {
		random := r.Intn(i + 1)
		table[random], table[i] = table[i], table[random]
	}
}

func printTable() (result string) {
	result += newLine
	for i, element := range table {
		if element == 0 {
			result += whiteBlock
		} else {
			result += fmt.Sprintf(" %02d", element)
		}

		if (i+1)%4 == 0 {
			result += newLine
		}
	}

	return
}

func startCaptureKeyBoard() {
	go func() {
		if err := termbox.Init(); err != nil {
			panic(err)
		}
		defer termbox.Close()

		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				{
					switch ev.Key {
					case termbox.KeyEsc:
						steps <- ESC
						break
					case termbox.KeyArrowUp:
						steps <- UP
						break
					case termbox.KeyArrowDown:
						steps <- DOWN
						break
					case termbox.KeyArrowLeft:
						steps <- LEFT
						break
					case termbox.KeyArrowRight:
						steps <- RIGHT
						break
					}
				}
			}
		}
	}()
}

func move(key enterKey) {

}

func run() {
	initTable()
	randomTable()
	startCaptureKeyBoard()
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

func main() {
	run()
	fmt.Println(printTable())
}
