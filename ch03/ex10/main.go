package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var b bytes.Buffer
	for i, v := range s {
		b.WriteRune(v)
		if i%3 == 2 {
			b.WriteRune(',')
		}
	}
	return b.String()
}

func main() {
	test := "12344543254623"
	fmt.Printf("%s -> %s\n", test, comma(test))
}
