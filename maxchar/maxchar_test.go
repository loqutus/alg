package main

import "testing"

func TestMaxChar(t *testing.T) {
	ans := maxChar("aabcde")
	if ans != rune('a') {
		t.Errorf("error %c", ans)
	}
}
