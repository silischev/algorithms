package main

import "log"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	s1 := "abcc"
	s2 := "cbca"

	log.Println(linear(s1, s2))
}

func linear(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1Map := map[string]int{}
	s2Map := map[string]int{}

	for _, l := range s1 {
		s1Map[string(l)]++
	}

	for _, l := range s2 {
		s2Map[string(l)]++
	}

	log.Println(s1Map)
	log.Println(s2Map)

	for _, l := range s1 {
		val, ok := s2Map[string(l)]

		if !ok {
			return false
		}

		if s1Map[string(l)] != val {
			return false
		}
	}

	return true
}
