package main

import "fmt"

func main() {
	A := []int{2, 2, 1, 3, 4, 2, 5, 2}
	fmt.Println(repeatedNTimes(A))
}

func repeatedNTimes(A []int) int {
    aimC := len(A) / 2
	for _, v := range A {
		count := 0
		for _, iv := range A {
			if v == iv {
				count++
			}
		}
		if count == aimC {
			return v
		}
	}
	return 0
}