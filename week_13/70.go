package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(45))
}

func climbStairs(n int) int {
    if n == 1 || n == 2 {
		return n
	}
	
	return climbStairs(n-1) + climbStairs(n-2)
	
}