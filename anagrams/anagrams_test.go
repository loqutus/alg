package main

import "testing"

func TestIsAnagram(t *testing.T) {
	ans := isAnagram("aab", "baa")
	if ans != true {
		t.Errorf("error: %t", ans)
	}
	ans = isAnagram("aac", "aab")
	if ans == true {
		t.Errorf("error: %t", ans)
	}
	ans = isAnagram("aac", "aaC")
	if ans != true {
		t.Errorf("error: %t", ans)
	}
	ans = isAnagram("aac!", "aaC")
	if ans != true {
		t.Errorf("error: %t", ans)
	}
}
