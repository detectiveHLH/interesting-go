package main

import (
	"fmt"
	"os"
	"time"
)

// 实现的🚀发射倒计时的程序，火箭会在5秒钟后发射
// 在倒数的过程中，如果输入任意的字符，回车之后就会终止发射
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