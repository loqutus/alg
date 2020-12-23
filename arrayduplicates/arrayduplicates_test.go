package main

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	input := []int{1, 1, 2, 3, 4, 4, 5}
	expected := []int{1, 2, 3, 4, 5}
	removeDuplicates(input)
	if len(input) != len(expected) {
		t.Errorf("error %d %d", len(input), len(expected))
		for _, v := range input {
			t.Errorf(fmt.Sprint(v))
		}
	}
}
