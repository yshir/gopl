package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string // 暗黙的に空文字 "" で初期化されている
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
