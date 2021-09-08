package main

import "testing"

func TestOneModification(t *testing.T) {
	cases := []struct {
		s      string
		result string
	}{
		{"aabcccccaaa", "a2b1c5a3"},
		{"a", "a"},
		{"ab", "ab"},
		{"aba", "aba"},
		{"abaa", "abaa"},
		{"abaaa", "abaaa"},
		{"abaaaa", "abaaaa"},
		{"abaaaaa", "a1b1a5"},
		{"abaaaaaaaaaa", "a1b1a10"},
		{"aab", "aab"},
		{"aabbbcc", "a2b3c2"},
		{"abbb", "abbb"},
		{"abbbc", "abbbc"},
	}

	for _, c := range cases {
		act := compressed(c.s)
		if act != c.result {
			t.Fatalf("Failed result [%s] for [%s], [%s]", act, c.s, c.result)
		}
	}
}
