package main

import "testing"

func TestReverseInt(t *testing.T) {
	ans := reverseInt(123)
	if ans != 321 {
		t.Errorf("error %d", ans)
	}
	ans = reverseInt(-123)
	if ans != -321 {
		t.Errorf("error %d", ans)
	}
}
