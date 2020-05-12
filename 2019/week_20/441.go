package main

import "fmt"

func main() {
	fmt.Println(arrangeCoins(100000000))
}

func arrangeCoins(n int) int {
    count := 0
	g := 1
	for n >= g{
		n -= g
		g++
		count++
	}
	
	return count
}