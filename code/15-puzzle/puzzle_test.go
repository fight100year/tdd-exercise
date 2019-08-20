package main

import (
	"testing"
)

func TestCreateTable(t *testing.T) {
	want := `
 01
`
	got := createTable()

	if want != got {
		t.Errorf("create table failed: want %v, got %v", want, got)
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
