package main

import (
	"fmt"
)

func anagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	h := map[rune]int{}
	for _, b1 := range s1 {
		if _, ok := h[b1]; ok {
			h[b1]++
		} else {
			h[b1] = 0
		}
	}
	for _, b2 := range s2 {
		if _, ok := h[b2]; ok {
			if h[b2] == 0 {
				delete(h, b2)
			} else {
				h[b2]--
			}
		} else {
			return false
		}
	}
	return len(h) == 0
}

func main() {
	s1 := "anagram"
	s2 := "anagram"

	fmt.Printf("%s and %s are anagram? -> %t\n", s1, s2, anagram(s1, s2))
}
