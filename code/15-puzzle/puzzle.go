package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

var (
	table = [4][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 0},
	}
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
}

func transform(i int) (int, int) {
	return i % 4, i / 4
}

func randomTable() {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	lastPos := 16
	for lastPos > 1 {
		changePos := r.Intn(lastPos)
		lastX, lastY := transform(lastPos - 1)
		changeX, changeY := transform(changePos)
		table[lastX][lastY], table[changeX][changeY] = table[changeX][changeY], table[lastX][lastY]

		lastPos--
	}
}

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
