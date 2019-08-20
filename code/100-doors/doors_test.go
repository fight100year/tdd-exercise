package main

import (
	"reflect"
	"testing"
)

func TestDoor(t *testing.T) {
	t.Run("config test", func(t *testing.T) {
		want := []int{1, 4, 9}
		got := configDoor(10)

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}

	})
}
