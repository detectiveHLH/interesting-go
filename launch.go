package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})

	go func() {
		_, _ = os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.")
	select {
	case <-time.After(5 * time.Second):
		break
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	fmt.Println("Launch.")
}