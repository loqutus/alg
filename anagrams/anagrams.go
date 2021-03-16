package main

import "unicode"

func isAnagram(stringA, stringB string) bool {
	mapA := make(map[rune]int)
	mapB := make(map[rune]int)
	for _, c := range stringA {
		c = unicode.ToLower(c)
		if unicode.IsLetter(c) {
			if val, ok := mapA[c]; ok {
				mapA[c] = val + 1
			} else {
				mapA[c] = 1
			}
		}
	}
	for _, c := range stringB {
		c = unicode.ToLower(c)
		if unicode.IsLetter(c) {
			if val, ok := mapB[c]; ok {
				mapB[c] = val + 1
			} else {
				mapB[c] = 1
			}
		}
	}
	if len(mapA) != len(mapB) {
		return false
	}
	for k, v := range mapA {
		if mapB[k] != v {
			return false
		}
	}
	return true
}
