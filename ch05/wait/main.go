package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// WaitForServer は URL のサーバへ接続を試みます。
// Exponential Backoff を使って1分間試みます。
// すべての試みが失敗したらエラーを報告します
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // Exponential Backoff
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func main() {
	url := "https://google.com"
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Success\n")
}
