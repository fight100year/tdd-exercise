package main

import (
	"math/rand"
	"time"
)

var (
	table = [4][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 0},
	}

	posX, posY int
)

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
