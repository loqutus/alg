package main

import "testing"

func TestPalindrome(t *testing.T) {
	ans := palindrome("abc")
	if ans != false {
		t.Errorf("error %t", ans)
	}
	ans = palindrome("abba")
	if ans != true {
		t.Errorf("error %t", ans)
	}
}
