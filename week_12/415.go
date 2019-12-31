package main

import (
	"fmt"
)

func main() {
	s1 := "9999"
	s2 := "9999"
	fmt.Println(addStrings(s1, s2))
}

func addStrings(num1 string, num2 string) string {
	num1I := make([]int, len(num1))
	num2I := make([]int, len(num2))
	for i, v := range num1 {
		num1I[i] = int(v)-48
	}
	for i, v := range num2 {
		num2I[i] = int(v)-48
	}
	adv := 0
	var resultArr []int
	if len(num1I) > len(num2I) {
		resultArr = exec(num1I, num2I, adv)
	} else {
		resultArr = exec(num2I, num1I, adv)
	}
	rs := make([]byte, 0)
	for _, v := range resultArr {
		rs = append(rs, byte(v+48))
	}
	return string(rs)
}

func exec(A, B []int, adv int) []int {
	for i := 0; i < len(B); i++ {
		if adv == 1 {
			if A[len(A)-1-i] + B[len(B)-1-i] + 1 > 9 {
				A[len(A)-1-i] = A[len(A)-1-i] + B[len(B)-1-i] + 1 - 10
			} else {
				A[len(A)-1-i] = A[len(A)-1-i] + B[len(B)-1-i] + 1
				adv = 0
			}
		} else {
			if A[len(A)-1-i] + B[len(B)-1-i] > 9 {
				A[len(A)-1-i] = A[len(A)-1-i] + B[len(B)-1-i] - 10
				adv = 1
			} else {
				A[len(A)-1-i] = A[len(A)-1-i] + B[len(B)-1-i]
			}
		}
	}
	if adv == 1 {
		for i := 0; i < len(A)-len(B); i++ {
			if adv == 1 {
				if A[len(A) - len(B) -i-1] + 1 > 9 {
					A[len(A) - len(B) -i-1] = A[len(A) - len(B) -i-1] + 1 - 10
				} else {
					A[len(A) - len(B) -i-1] = A[len(A) - len(B) -i-1] + 1
					adv = 0
				}
			}
		}
	}
	if adv == 1 {
		rArr := make([]int, len(A)+1)
		rArr[0] = 1
		for i, v := range A {
			rArr[i+1] = v
		}
		return rArr
	} 
	return A
}