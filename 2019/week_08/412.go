package main

import (
	"fmt"
	"strconv"
	)

func main() {
	fmt.Println(fizzBuzz(15))
}


func fizzBuzz(n int) []string {
	fbArr := make([]string, 0)
	for i := 1; i <= n; i++ {
		if i % 15 == 0 {
			fbArr = append(fbArr, "FizzBuzz")
			continue
		}
		if i % 5 == 0 {
			fbArr = append(fbArr, "Buzz")
			continue
		}
		if i % 3 == 0 {
			fbArr = append(fbArr, "Fizz")
			continue
		}
		fbArr = append(fbArr, strconv.Itoa(i))
    }
	return fbArr
}