package main

func maxChar(s string) rune {
	d := make(map[rune]int)
	rs := []rune(s)
	for _, c := range rs {
		if val, ok := d[c]; ok {
			d[c] = val + 1
		} else {
			d[c] = 1
		}
	}
	var a rune
	mN := 1
	for c, n := range d {
		if n > mN {
			a = rune(c)
			mN = n
		}
	}
	return a
}
