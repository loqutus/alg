package main

import (
	"fmt"
	"math"
)

func pyramid(n int) {
	midpoint := int(math.Floor(float64((2*n - 1) / 2)))
	for row := 0; row < n; row++ {
		level := ""
		for column := 0; column < 2*n-1; column++ {
			if midpoint-row <= column && midpoint+row >= column {
				level += "#"
			} else {
				level += " "
			}
		}
		fmt.Println(level)
	}
}

func main() {
	pyramid(5)
}
