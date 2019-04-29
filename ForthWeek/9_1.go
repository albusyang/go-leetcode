package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(10022201))
}

func isPalindrome(x int) bool {
	var t int
	for a := x; a > 0; a /= 10 {
		r := a % 10
		t = t*10 + r
	}
	return (t == x)
}
