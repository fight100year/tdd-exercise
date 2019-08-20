package main

import (
	"testing"
)

func TestCreateTable(t *testing.T) {
	want := ` 01
`
	got := createTable()

	if want != got {
		t.Errorf("create table failed: want %v, got %v", want, got)
	}
}
