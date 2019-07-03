package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(4))
}

// Fibonacci: 0,1,1,2,3,5,8,13,21...
func climbStairs(n int) int {
    if n == 1 || n == 2 {
		return n
	}
	nn := 1
	temp := 0
	for n != 0 {
		nn += temp
		temp = nn - temp
		n--
	}
	return nn
}