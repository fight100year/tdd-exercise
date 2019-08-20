package main

import (
	"testing"
)

func TestInitTable(t *testing.T) {
	initTable()

	want := [16]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	if want != table {
		t.Errorf("create table failed: want %v, got %v", want, table)
	}
}

func TestRandomTable(t *testing.T) {
	randomTable()

	m := make(map[int]bool, 16)
	for _, element := range table {
		m[element] = true
	}

	if len(m) != 16 {
		t.Errorf("excepted 16 unique number, got %v", table)
	}
}

func TestPrintTable(t *testing.T) {
	table = [16]int{4, 3, 2, 1, 8, 7, 6, 5, 12, 11, 10, 9, 0, 15, 14, 13}
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
