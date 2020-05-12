package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(28))
}

func convertToTitle(n int) string {
	cStr := ""
	
	for n > 0 {
		t := (n-1)%26
		s := string(t+65)
		cStr = s + cStr
		n = (n-1)/26
	}
	
	return cStr
}