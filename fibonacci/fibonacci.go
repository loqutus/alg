package main

var d map[int]int

func fibonacci(f int) int {
	if f < 2 {
		return f
	}
	if val, ok := d[f]; ok {
		return val
	} else {
		val = fibonacci(f-1) + fibonacci(f-2)
		d[f] = val
		return val
	}
}

func init() {
	d = make(map[int]int)

}
