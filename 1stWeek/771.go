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
    for _, jv := range J {
		for _, sv := range S {
			if jv == sv {
				jSum++
			}
		}
	}
	return jSum
}