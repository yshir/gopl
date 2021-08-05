package main

import (
	"fmt"
)

func basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	tests := []struct {
		input    string
		expected string
	}{
		{"a", "a"},
		{"a.go", "a"},
		{"a/b/c.go", "c"},
		{"a/b.c.go", "b.c"},
	}

	for _, tt := range tests {
		if basename(tt.input) != tt.expected {
			fmt.Printf("basename return wrong string. want=%s, got=%s\n",
				tt.expected, basename(tt.input))
		}
	}
}
