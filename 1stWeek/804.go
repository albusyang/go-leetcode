package main

import (
	"fmt"
)

func main() {

	var ss = make([]string, 0)
	ss = append(ss, "gin")
	ss = append(ss, "zen")
	ss = append(ss, "gig")
	ss = append(ss, "msg")
	ss = append(ss, "ggg")
	ss = append(ss, "qq")
	fmt.Println(uniqueMorseRepresentations(ss))
}

func uniqueMorseRepresentations(words []string) int {
	var morseCode = [...]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	mIndex := 97
	morseMap := make(map[int]string)
	for _, v := range morseCode {
		morseMap[mIndex] = v
		mIndex++
	}

	wMorse := make([]string, 0)

	for _, vw := range words {
		inMorse := ""
		for _, vs := range vw {
			inMorse = inMorse + morseMap[int(vs)]
		}

		if len(wMorse) == 0 {
			wMorse = append(wMorse, inMorse)
		} else {
			flag := false
			for _, vwm := range wMorse {
				if vwm == inMorse {
					flag = true
				}
			}
			if flag == false {
				wMorse = append(wMorse, inMorse)
			}
		}
	}
	return len(wMorse)
}
