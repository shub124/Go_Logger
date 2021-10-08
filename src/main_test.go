package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	if Calculate(2) != 5 {
		t.Error("2 + 3 should be equal to 4")
	}
}
