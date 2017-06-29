package main

import "testing"

func TestPlus(t *testing.T) {
	t.Run("Good", func(t *testing.T) {
		a := Plus(10, 15)
		if a != 25 {
			t.Error(a)
		}
	})

	t.Run("Error", func(t *testing.T) {
		a := Plus(10, 15)
		if a != 26 {
			t.Error(a)
		}
	})
}
