package main

func matrix(n int) [][]int {
	results := make([][]int, n)
	for i := 0; i < n; i++ {
		results[i] = make([]int, n)
	}
	counter := 1
	startColumn := 0
	endColumn := n - 1
	startRow := 0
	endRow := n - 1
	for startColumn <= endColumn && startRow <= endRow {
		// top row
		for i := startColumn; i <= endColumn; i++ {
			results[startRow][i] = counter
			counter++
		}
		startRow++
		// right column
		for i := startRow; i <= endRow; i++ {
			results[i][endColumn] = counter
			counter++
		}
		endColumn--
		// bottom row
		for i := endColumn; i >= startColumn; i-- {
			results[endRow][i] = counter
			counter++
		}
		endRow--
		// start column
		for i := endRow; i >= startRow; i-- {
			results[i][startColumn] = counter
			counter++
		}
		startColumn++
	}
	return results
}
