package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // "/" が見つからなければ -1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	test := "/foo/bar/baz.jpg"
	fmt.Printf("%s -> %s\n", test, basename(test))
}
