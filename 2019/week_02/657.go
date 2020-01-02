package main

import (
	"fmt"
)

func main() {
	fmt.Println(judgeCircle("UUDDRRLL"))
}

func judgeCircle(moves string) bool {

	countUD := 0
	countRL := 0

	for _, v := range moves {
		switch v {
		case 'U':
			countUD++
		case 'D':
			countUD--
		case 'R':
			countRL++
		case 'L':
			countRL--
		}
	}

	if countUD == 0 && countRL == 0 {
		return true
	}
	return false
}
