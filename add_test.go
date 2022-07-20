package main

import (
	"testing"
)

func TestAdd(t *testing.T) {

	t.Run("Postive number", func(t *testing.T) {
		got := AddPositiveNumber(4, 5, 6)
		expected := 15

		if got != expected {
			t.Errorf("got: %d expected %d", got, expected)
		}
	})
	t.Run("Negative number", func(t *testing.T) {
		got := AddPositiveNumber(-4, 5, 6)
		expected := -1

		if got != expected {
			t.Errorf("got: %d expected %d", got, expected)
		}
	})
	t.Run("Postive number", func(t *testing.T) {
		got := AddPositiveNumber(99999, 5, 1)
		expected := 100005

		if got != expected {
			t.Errorf("got: %d expected %d", got, expected)
		}
	})

}
