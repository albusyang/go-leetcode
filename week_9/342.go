package main

import "fmt"

func main() {
	fmt.Println(isPowerOfFour(0))
}


func isPowerOfFour(num int) bool {
	if num == 1 || num == 4{
		return true
	}
	
	n := divideFour(num)
	for true {
		if divideFour(n) == 0 {
			return false
		}
		if divideFour(n) == 1 {
			return true
		}
		n = divideFour(n)
	}
    return false
}

func divideFour(n int) int {
	if n % 4 != 0 {
		return 0
	}
	return n / 4
}