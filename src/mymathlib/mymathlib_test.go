package mymathlib

import "testing"

func TestMyMathLib(t *testing.T) {
	// initialization
	a := 1
	b := 2
	// execution
	sum := add(a, b)
	// validation
	if sum != 3 {
		t.Errorf("Expected 3 , found = %v!", sum)
	}
}
