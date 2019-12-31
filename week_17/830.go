package main

import "fmt"

func main() {
	fmt.Println(largeGroupPositions("abbxxx"))
}

func largeGroupPositions(S string) [][]int {
	var temp rune
	var start, finish int
	ss := make([][]int, 0)
	for i, v := range S {
		if v != temp {
			temp = v
			if finish - start > 1 {
				sf := make([]int, 2)
				sf[0] = start
				sf[1] = finish
				ss = append(ss, sf)
			}
			start = i
			finish = i
		} else {
			finish++
		}
	}
	
	if finish - start > 1 {
		sf := make([]int, 2)
		sf[0] = start
		sf[1] = finish
		ss = append(ss, sf)
	}
    
	return ss
}