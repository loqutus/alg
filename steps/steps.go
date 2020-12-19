package main

import (
	"fmt"
)

func steps(n int) {
	s := "#"
	for i := 0; i < n; i++ {
		s += " "
	}
	fmt.Println(s)
	for i := 1; i < n; i++ {
		x := []byte(s)
		x[i] = byte('#')
		s = string(x)
		fmt.Println(s)
	}
	return
}
