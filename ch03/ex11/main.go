package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var b bytes.Buffer
	if s[0] == '+' || s[0] == '-' {
		b.WriteByte(s[0])
		s = s[1:]
	}
	for i, v := range s {
		b.WriteRune(v)
		if i%3 == 2 {
			b.WriteRune(',')
		}
	}
	return b.String()
}

func main() {
	test := "+12344543254623"
	fmt.Printf("%s -> %s\n", test, comma(test))
}
