package main

/*
([)]
(){]
()[]
()[]{}
[{}]
*/
func isValid(s string) bool {
	brackets := []string{}
	bracketsMap := map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}

	for _, b := range s {
		bracket := string(b)

		if isOpenBracket(bracket) {
			brackets = append(brackets, bracket)
			continue
		}

		if len(brackets) == 0 {
			return false
		}

		last := brackets[len(brackets)-1]
		if bracketsMap[last] != bracket {
			return false
		}

		brackets = brackets[:len(brackets)-1]
	}

	return len(brackets) == 0
}

func isOpenBracket(bracket string) bool {
	for _, s := range []string{"(", "{", "["} {
		if bracket == s {
			return true
		}
	}

	return false
}
