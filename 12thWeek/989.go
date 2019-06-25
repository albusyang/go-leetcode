package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr1 := []int{9,9,9,9,9,9,9,9,9,9}
	kk := 999
	fmt.Println(addToArrayForm(arr1, kk))
}

func addToArrayForm(A []int, K int) []int {
	ks := strconv.Itoa(K)
	kArr := make([]int, len(ks))
	for i, v := range ks {
		kArr[i] = int(v)-48
	}
	adv := 0
	if len(A) > len(kArr) {
		return exec(A, kArr, adv)
	} else {
		return exec(kArr, A, adv)
	}
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