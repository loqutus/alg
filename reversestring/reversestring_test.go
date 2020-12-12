package main

import "testing"

func TestReverseString(t *testing.T) {
	ans := reverseString("abc")
	if ans != "cba" {
		t.Errorf("error %s", ans)
	}
}
