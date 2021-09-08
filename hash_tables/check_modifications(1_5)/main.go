package main

import (
	"log"
	"math"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	s1 := "bake"
	s2 := "pale"

	log.Println(isOneModification(s1, s2))
}

func isOneModification(s1, s2 string) bool {
	if len(s1)-len(s2) > int(math.Abs(1.0)) {
		return false
	}

	if len(s1) == 1 && len(s2) == 1 && s1[0] != s2[0] {
		return false
	}

	hasDiff := false
	i := 0
	for i <= len(s1) {
		if i >= len(s1) || i >= len(s2) {
			return true
		}

		if s1[i] != s2[i] {
			if len(s1) == len(s2) {
				if s1[i+1] != s2[i+1] || hasDiff {
					return false
				}

				hasDiff = true
			}

			if len(s1) != len(s2) {
				if len(s1) > len(s2) && s1[i+1] != s2[i] {
					return false
				} else if len(s1) < len(s2) && s1[i] != s2[i+1] {
					return false
				}

				i++
			}
		}

		i++
	}

	return true
}
