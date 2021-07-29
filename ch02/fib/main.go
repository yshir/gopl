package main

import "fmt"

func main() {
	for i := 0; i < 20; i++ {
		fmt.Printf("n = %d -> fib = %d\n", i, fib(i))
	}
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
