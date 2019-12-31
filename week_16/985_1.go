package main

import "fmt"

func main() {
	A := []int{1,2,3,4}
	queries := [][]int{{1,0}, {-3,1}, {-4,0}, {2,3}}
	fmt.Println(sumEvenAfterQueries(A, queries))
}

func sumEvenAfterQueries(A []int, queries [][]int) []int {
    result := make([]int, len(queries))
	sum := 0
	for _, v := range A {
		if v&1 == 0 {
			sum += v
		}
	}
	for i := 0; i < len(queries); i++ {
		r := queries[i]
		ai := A[r[1]]
		A[r[1]] += r[0]
		
		if ai&1 == 0 {
			sum -= ai
		}
		if A[r[1]]&1 == 0 {
			sum = sum + A[r[1]]
		}
		result[i] = sum
	}
	return result
}