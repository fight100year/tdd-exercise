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

	posX, posY int
	countStep  int
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
	return i / 4, i % 4
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

	findPos()
}

func findPos() {
	for x := range table {
		for y := range table[x] {
			if table[x][y] == 0 {
				posX, posY = x, y
			}
		}
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

func show() {
	fmt.Printf("total step: %d\n", countStep)
	countStep++
	fmt.Println(printTable())
}

func move(key enterKey) {
	switch key {
	case UP:
		if posX != 0 {
			table[posX][posY], table[posX-1][posY] = table[posX-1][posY], table[posX][posY]
			posX--
			show()
		}
		break
	case DOWN:
		if posX != 3 {
			table[posX][posY], table[posX+1][posY] = table[posX+1][posY], table[posX][posY]
			posX++
			show()
		}
		break
	case LEFT:
		if posY != 0 {
			table[posX][posY], table[posX][posY-1] = table[posX][posY-1], table[posX][posY]
			posY--
			show()
		}
		break
	case RIGHT:
		if posY != 3 {
			table[posX][posY], table[posX][posY+1] = table[posX][posY+1], table[posX][posY]
			posY++
			show()
		}
		break
	}
}

func run() {
	fmt.Println(1)
	initTable()
	randomTable()
	show()
	startCaptureKeyBoard()
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
