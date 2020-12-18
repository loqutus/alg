package main

import (
	"strings"
	"unicode"
)

func capitalize(str string) string {
	split := strings.Split(str, " ")
	x := split[0]
	result := string(unicode.ToUpper(rune(x[0])))
	result += split[0][1:]
	for _, s := range strings.Split(str, " ")[1:] {
		result += " "
		result += string(unicode.ToUpper(rune(s[0])))
		result += s[1:]
	}
	return result
}
