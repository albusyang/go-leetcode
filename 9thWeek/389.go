package main

import (
    "fmt"
)

func main() {
    fmt.Println(findTheDifference("abcd", "abcde"))
}

func findTheDifference(s string, t string) byte {
    var is, it rune
    for _, v := range s {
        is += v
    }
    for _, v:= range t {
        it += v
    }
    return byte(it-is)
}
