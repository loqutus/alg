package main

import "testing"

func TestCapitalize(t *testing.T) {
	input := "hello, world!"
	result := capitalize(input)
	if result != "Hello, World!" {
		t.Errorf("error %s", result)
	}
}
