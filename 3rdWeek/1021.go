package main

import "fmt"

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
}

func removeOuterParentheses(S string) string {
    l := 0
	r := 0
	result := make([]rune, 0)
	for _, v := range S {
		if v == '(' {
			l++
		}
		if v == ')' {
			r++
		}
		if l + r > 1 && l != r {
			result = append(result, v)
		}
		if l == r {
			l = 0
			r = 0
		}
	}
	return string(result)
}