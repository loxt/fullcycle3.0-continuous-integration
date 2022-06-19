package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(10, 10)
	if total != 20 {
		t.Errorf("Sum result is invalid: %d", total)
	}
}
