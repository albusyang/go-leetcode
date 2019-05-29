package main

import "fmt"

func main() {
	fmt.Println(isPowerOfThree(27))
}


func isPowerOfThree(n int) bool {
	if n == 1 || n == 3{
		return true
	}
	if n % 2 == 0 {
		return false
	}
	num := divideThree(n)
	for true {
		if divideThree(num) == 0 {
			return false
		}
		if divideThree(num) == 1 {
			return true
		}
		num = divideThree(num)
	}
    return false
}

func divideThree(n int) int {
	if n % 3 != 0 {
		return 0
	}
	return n / 3
}