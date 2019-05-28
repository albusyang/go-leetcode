package main

import "fmt"

func main() {
	words := []string{"app", "well", "mini", "pit"}
	fmt.Println(findWords(words))
}

func findWords(words []string) []string {
	if len(words) == 1 {
		if words[0] == "" {
			return words
		}
	}
	oneArr := "QWERTYUIOPqwertyuiop"
	twoArr := "ASDFGHJKLasdfghjkl"
	thrArr := "ZXCVBNMzxcvbnm"
	
	lineWords := make([]string, 0)
	
	for _, v := range words {
		oneCount := 0
		twoCount := 0
		thrCount := 0
		for _, inV := range v{
			
			for _, one := range oneArr{
				if inV == one {
					oneCount++
				}
			}
			if oneCount == len(v) {
				lineWords = append(lineWords, v)
				break
			}
			
			for _, two := range twoArr{
				if inV == two {
					twoCount++
				}
			}
			if twoCount == len(v) {
				lineWords = append(lineWords, v)
				break
			}
			
			for _, thr := range thrArr{
				if inV == thr {
					thrCount++
				}
			}
			if thrCount == len(v) {
				lineWords = append(lineWords, v)
				break
			}
		}
	}
    return lineWords
}
