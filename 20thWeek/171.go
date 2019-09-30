package main

import "fmt"

func main() {
	fmt.Println(titleToNumber("ZY"))
}

func titleToNumber(s string) int {
	cNum := 0
	for i := 0; i < len(s); i++ {
		temp := 1
		for j := 1; j < len(s) - i; j++ {
			temp *= 26
		}
		cNum += (int(s[i]) - 64) * temp
	}
	
	return cNum
}