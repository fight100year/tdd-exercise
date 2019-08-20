package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	table = [16]int{}
)

const (
	whiteBlock = "   "
	newLine    = "\n"
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

func main() {
	initTable()
	randomTable()
	fmt.Println(printTable())
}
