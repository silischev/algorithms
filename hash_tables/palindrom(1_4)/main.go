package main

import "log"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	s := "appa" // apzpa appzza

	log.Println(s, len(s))
	log.Println(isPalindrom(s))
}

func isPalindrom(s string) bool {
	letterCnt := map[string]int{}
	sLen := len(s)

	for _, l := range s {
		letterCnt[string(l)]++
	}

	oddLettersCnt := 0

	for _, v := range letterCnt {
		if sLen%2 != 0 {
			if oddLettersCnt > 1 || (v%2 != 0 && v > 1) {
				return false
			}

			if v == 1 {
				oddLettersCnt++
			}
		}

		if sLen%2 == 0 && v%2 != 0 {
			return false
		}
	}

	return true
}
