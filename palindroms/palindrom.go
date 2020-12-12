package main

func palindrome(s string) bool {
	l := len(s)
	j := l - 1
	i := 0
	for i < l/2 && j > l/2 {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
