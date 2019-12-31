package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
    sx := strconv.FormatInt(int64(x), 2)
	sy := strconv.FormatInt(int64(y), 2)
	sxLen := len(sx)
	syLen := len(sy)
	for i := 0; i < 32 - sxLen; i++ {
		sx = "0" + sx
	}
	for i := 0; i < 32 - syLen; i++ {
		sy = "0" + sy
	}
	hamming := 0
	for i := 0; i < 32; i++ {
		if sx[i] != sy[i] {
			hamming = hamming + 1
		}
	}
	return hamming
}