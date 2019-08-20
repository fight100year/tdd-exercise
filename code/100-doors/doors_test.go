package main

import (
	"reflect"
	"testing"
)

func TestDoor(t *testing.T) {
	t.Run("config test", func(t *testing.T) {
		want := []int{1, 0, 0, 1, 0}
		got := configDoor(5)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}

	})
}
