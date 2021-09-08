package main

import "testing"

func TestOneModification(t *testing.T) {
	cases := []struct {
		s1     string
		s2     string
		result bool
	}{
		{"apple", "aple", true},
		{"apple", "apzle", true},
		{"pale", "pale", true},
		{"pale", "ple", true},
		{"ple", "pale", true},
		{"pales", "pale", true},
		{"pales", "bale", false},
		{"pale", "pales", true},
		{"pale", "bale", true},
		{"bale", "pale", true},
		{"bake", "pale", false},
		{"pale", "bake", false},
		{"z", "b", false},
		{"zx", "b", false},
		{"zx", "bx", true},
		{"zx", "bz", false},
		{"zx", "bxz", false},
		{"zx", "bzx", true},
		{"zx", "azx", true},
		{"zx", "az", false},
	}

	for _, c := range cases {
		act := isOneModification(c.s1, c.s2)
		if act != c.result {
			t.Fatalf("Failed result [%t] for [%s], [%s]", act, c.s1, c.s2)
		}
	}
}
