package main

import (
	"fmt"
)

func main() {
	fmt.Println(toLowerCase("Hello&HI"))
}

func toLowerCase(str string) string {
	lstr := ""
    for _, v := range str {
		if v < 91 && v > 64{
			v += 32
		}
		lstr = lstr + string(v)
	}
	return lstr
} 