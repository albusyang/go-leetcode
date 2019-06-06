package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
	nn := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(nn)
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}
    lnl := make([]int, 0)
	for _, v := range s {
		if v >= 48 && v <= 57 {
			lnl = append(lnl, int(v))
		}
		if v >= 65 && v <= 90 {
			lnl = append(lnl, int(v))
		}
		if v >= 97 && v <= 122 {
			lnl = append(lnl, int(v) - 32)
		}
	}
	if len(lnl) == 1 {
		return true
	}
	l := 0
	r := 0
	isP := true
	if len(lnl) % 2 == 0 {
		l = len(lnl)/2 - 1
		r = len(lnl)/2
	} else {
		l = len(lnl)/2 - 1
		r = len(lnl)/2 + 1
	}
	for i := 0; i < len(lnl)/2; i++ {
		if lnl[l-i] != lnl[r+i] {
				isP = false
				break
		}			
	}
	return isP
}