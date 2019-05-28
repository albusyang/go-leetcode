package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(64))
}


func isPowerOfTwo(n int) bool {
	if n == 1 || n == 2{
		return true
	}
	if n % 2 == 1 {
		return false
	}
	num := divideTwo(n)
	for true {
		if divideTwo(num) == 0 {
			return false
		}
		if divideTwo(num) == 1 {
			return true
		}
		num = divideTwo(num)
	}
    return false
}

func divideTwo(n int) int {
	if n % 2 == 1 {
		return 0
	}
	return n / 2
}