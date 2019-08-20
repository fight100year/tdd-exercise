package main

import (
	"reflect"
	"testing"
)

func TestDoor(t *testing.T) {
	t.Run("config test", func(t *testing.T) {
		var doors Doors
		doors.setDoorCount(10)
		got := doors.Do()
		want := []int{1, 4, 9}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %v, got %v", want, got)
		}

	})
}
