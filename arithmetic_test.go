package main

import "testing"

func TestSum(t *testing.T) {
	sum := Sum(2, 3)
	if sum != 5 {
		t.Errorf("Sum was incorrect, got: %d, expected: %d.", sum, 5)
	}
}
