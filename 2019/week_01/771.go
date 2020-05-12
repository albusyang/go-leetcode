package main

import (
	"fmt"
)

func main() {

	jSum := numJewelsInStones("ab", "aaabcabc")
	fmt.Println(jSum)
}

func numJewelsInStones(J string, S string) int {
	jSum := 0
	for _, j := range J {
		for _, s := range S {
			if j == s {
				jSum++
			}
		}
	}
	return jSum
}
