package main

import "fmt"

func main() {
	s := []byte{'h','e','l','l','o'}
	fmt.Println(s)
	reverseString(s)
	fmt.Println(string(s))
}

func reverseString(s []byte)  {
    for i := 0; i < len(s) / 2; i++ {
		s[i], s[len(s) - i - 1] = s[len(s) - i - 1], s[i]
	}
}