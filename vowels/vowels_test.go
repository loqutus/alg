package main

import (
	"testing"
)

func TestVowels(t *testing.T) {
	input := "Hi There!"
	output := vowels(input)
	if output != 3 {
		t.Errorf("error: %d", output)
	}
}
