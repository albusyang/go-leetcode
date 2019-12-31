package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(10022201))
}

func isPalindrome(x int) bool {
	//arrOne := make([]int, 0)
	if x < 0 {
		return false
	}
	begin := false
	arrOne := []int{}
	//haFlag := true
	arrTwo, haFlag:= getNum(arrOne, 1000000000000000000, x, begin)
	fmt.Println("arrTwo: ", arrTwo)
	if !haFlag {
		return haFlag
	}
	
	for i := 0; i < len(arrTwo); i++ {
		
		if arrTwo[i] != arrTwo[len(arrTwo) - 1 - i] {
			haFlag = false
			break
		}
	}
	return haFlag
}

func getNum(arr []int, nums, ss int, begin bool) ([]int, bool) {
	fmt.Println("arr: ", arr)
	if nums == 0 {
		return arr, true
	}
	if ss >= nums {
		if nums < 10 {
			arr = append(arr, ss)
			return arr, true
		}
		begin = true
		v := ss / nums
		ssNext := ss % nums
		if ssNext == 0 {
			return arr, false
		}
		arr = append(arr, v)
		return getNum(arr, nums/10, ssNext, begin)
	} else {
		if begin {
			arr = append(arr, 0)
		}
		return getNum(arr, nums/10, ss, begin)
	}
}
