package main

import "fmt"

func main() {
	order := "DIDI"
	fmt.Println(diStringMatch(order))
}

func diStringMatch(S string) []int {
    haka := make([]int,len(S)+1)
	l := 0
	r := len(S)
	for i := 0; i < len(S); i++ {
		if S[i] == 73 {
			haka[i] = l
			l++
		} else {
			haka[i] = r
			r--
		}
	}
	haka[len(S)] = l
	return haka
}