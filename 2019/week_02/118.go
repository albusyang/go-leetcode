package main

import "fmt"

func main() {
	fmt.Println(generate(50))
}

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{{1}}
	}

	result := [][]int{{1}, {1, 1}}
	times := numRows - 2
	head := 1

	for times >= 1 {
		newRow := nextRow(result[head])
		result = append(result, newRow)
		times--
		head++
	}
	return result
}

func nextRow(row []int) []int {
	newRow := make([]int, len(row)+1)
	newRow[0], newRow[len(row)] = 1, 1
	
	for i := 1; i < len(row); i++ {
		newRow[i] = row[i-1] + row[i]
	}

	return newRow
}
