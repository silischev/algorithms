package main

import (
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	s1 := "aabcccccaaa"

	log.Println(compressed(s1))
}

func compressed(s string) string {
	sLen := len(s)
	seqLetterCnt := 1
	res := ""

	if sLen <= 2 {
		return s
	}

	for i := 0; i < sLen; i++ {
		l := string(s[i])

		if i < sLen-1 && s[i] == s[i+1] {
			seqLetterCnt++
			continue
		}

		res += fmt.Sprintf("%s%v", l, seqLetterCnt)
		seqLetterCnt = 1
	}

	if sLen <= len(res) {
		return s
	}

	return res
}
