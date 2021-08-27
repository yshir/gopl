package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	arr := []string{"a", "b", "", "", "c"}
	fmt.Printf("arr = %q\n", arr)
	fmt.Printf("nonempty(arr) = %q\n", nonempty(arr))
}
