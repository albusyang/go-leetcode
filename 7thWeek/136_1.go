package main

import "fmt"

func main() {
	nums := []int{4,1,2,1,2}
	fmt.Println(singleNumber(nums))
}

func singleNumber(nums []int) int {
    n := 0
    for _,v:=range nums{
        n ^=v
	}
	return n
}