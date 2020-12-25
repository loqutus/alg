package main

import "unicode"

func vowels(str string) int {
	v := []rune{'a', 'e', 'i', 'o', 'u'}
	var s int
	for _, c := range str {
		found := false
		for _, w := range v {
			if unicode.ToLower(c) == w {
				found = true
				break
			}
		}
		if found {
			s++
		}
	}
	return s
}
