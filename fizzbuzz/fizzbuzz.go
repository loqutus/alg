package main

import "fmt"

func fizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("fizzbuzz")
			} else {
				fmt.Println("fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}
