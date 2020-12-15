package main

func chunk(arr []int, chunkSize int) [][]int {
	var res [][]int
	currentElem := 0
	currentChunk := 0
	for _, e := range arr {
		if currentElem < chunkSize {
			res[currentChunk][currentElem] = e
			currentElem++
		} else {
			currentElem = 0
			currentChunk++
		}
	}
	return res
}
