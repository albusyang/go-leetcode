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
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	pt := [][]int{{1}, {1, 1}}
	n := numRows - 2
	rr := 1

	for n >= 1 {
		nr := nextRow(pt[rr])
		pt = append(pt, nr)
		n--
		rr++
	}
	return pt
}

func nextRow(row []int) []int {
	nr := make([]int, len(row)+1)
	nr[0], nr[len(row)] = 1, 1
	for i := 1; i < len(row); i++ {
		nr[i] = row[i-1] + row[i]
	}
	return nr
}
