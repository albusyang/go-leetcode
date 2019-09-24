package main

import (
	"fmt"
)

func main() {
	fmt.Println(defangIPaddr("255.100.50.0"))
}

func defangIPaddr(address string) string {
	result := ""
	symbol := "[.]"
	for _, v := range address {
		vs := string(v)
		if vs == "." {
			result += symbol
			continue
		}
		result += vs
	}
	return result
}