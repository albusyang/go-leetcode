package main

import "fmt"

func main() {
	A := []int{1,2,3,4}
	queries := [][]int{{1,0}, {-3,1}, {-4,0}, {2,3}}
	fmt.Println(sumEvenAfterQueries(A, queries))
}

func sumEvenAfterQueries(A []int, queries [][]int) []int {
    result := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		r := queries[i]
		index := r[1]
		A[index] += r[0]
		
		oSum := 0
		for _, v := range A {
			if v&1 == 0 {
				oSum += v
			}
		}
		result[i] = oSum
	}
	return result
}