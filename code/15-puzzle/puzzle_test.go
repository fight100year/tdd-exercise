package main

import (
	"testing"
)

func TestRandomTable(t *testing.T) {
	randomTable()

	m := make(map[int]bool, 16)
	for x := range table {
		for y := range table[x] {
			m[table[x][y]] = true

		}
	}

	if len(m) != 16 {
		t.Errorf("excepted 16 unique number, got %v", table)
	}
}

func TestPrintTable(t *testing.T) {
	table = [4][4]int{
		{4, 3, 2, 1},
		{8, 7, 6, 5},
		{12, 11, 10, 9},
		{0, 15, 14, 13},
	}
	want := `
 04 03 02 01
 08 07 06 05
 12 11 10 09
    15 14 13
`
	got := printTable()

	if want != got {
		t.Errorf("create table failed: want %v, got %v", want, got)
	}
}

func TestMove(t *testing.T) {
	cases := []struct {
		name  string
		input [4][4]int
		want  [4][4]int
		key   enterKey
	}{
		{
			"accept scenes",
			[4][4]int{
				{4, 3, 2, 1},
				{8, 7, 6, 5},
				{12, 11, 10, 9},
				{0, 15, 14, 13},
			},

			[4][4]int{
				{4, 3, 2, 1},
				{8, 7, 6, 5},
				{0, 11, 10, 9},
				{12, 15, 14, 13},
			},
			UP,
		},
		{
			"accept scenes",
			[4][4]int{
				{4, 3, 2, 1},
				{8, 7, 6, 5},
				{12, 11, 10, 9},
				{0, 15, 14, 13},
			},

			[4][4]int{
				{4, 3, 2, 1},
				{8, 7, 6, 5},
				{12, 11, 10, 9},
				{15, 0, 14, 13},
			},
			RIGHT,
		},
		{
			"unaccept scenes",
			[4][4]int{
				{4, 3, 2, 1},
				{8, 7, 6, 5},
				{12, 11, 10, 9},
				{0, 15, 14, 13},
			},

			[4][4]int{
				{4, 3, 2, 1},
				{8, 7, 6, 5},
				{12, 11, 10, 9},
				{0, 15, 14, 13},
			},
			LEFT,
		},
		{
			"unaccept scenes",
			[4][4]int{
				{4, 3, 2, 1},
				{8, 7, 6, 5},
				{12, 11, 10, 9},
				{0, 15, 14, 13},
			},

			[4][4]int{
				{4, 3, 2, 1},
				{8, 7, 6, 5},
				{12, 11, 10, 9},
				{0, 15, 14, 13},
			},
			DOWN,
		},
	}

	for _, x := range cases {
		t.Run(x.name, func(t *testing.T) {
			table = x.input
			findPos()
			move(x.key)
			key := func(k enterKey) string {
				switch k {
				case UP:
					return "UP"
				case DOWN:
					return "DOWN"
				case LEFT:
					return "LEFT"
				case RIGHT:
					return "RIGHT"
				default:
					return "unknown key"
				}

			}(x.key)

			if table != x.want {
				t.Errorf("failed, key:%s want:%v got:%v", key, x.want, table)
			}
		})
	}
}
