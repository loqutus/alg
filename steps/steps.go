package main

import "fmt"

func steps(n int) {
	s := "#"
	for i := 0; i < n; i++ {
		fmt.Println(s)
		s += "#"
	}
	return
}
