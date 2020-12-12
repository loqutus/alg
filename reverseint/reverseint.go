package main

import (
	"strconv"
)

func reverseInt(s int) int {
	x := strconv.Itoa(s)
	var r string
	var p int
	if s < 0 {
		p = 1
	} else {
		p = 0
	}
	for i := len(x) - 1; i >= p; i-- {
		r += string(x[i])
	}
	y, _ := strconv.Atoi(r)
	if s < 0 {
		return -y
	} else {
		return y
	}
}
