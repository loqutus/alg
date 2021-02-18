package main

import (
	"testing"
)

func TestPyramid(t *testing.T) {
	output := fibonacci(7)
	if output != 13 {
		t.Errorf("error: %d", output)
	}
}
