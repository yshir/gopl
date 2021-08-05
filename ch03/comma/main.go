package main

import "fmt"

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	test := "123456789"
	fmt.Printf("%s -> %s\n", test, comma(test))
}
