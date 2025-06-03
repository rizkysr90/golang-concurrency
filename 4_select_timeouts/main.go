// Practice using select with multiple channels and a timeout.
package main

import (
	"fmt"
	"time"
)

func slowProcess(ch chan<- string) {
	time.Sleep(time.Second * 2)
	ch <- "Result from slow process"
}

func fastProcess(ch chan<- string) {
	time.Sleep(time.Second * 1)
	ch <- "Result from fast process"
}

func main() {
	slow := make(chan string)
	fast := make(chan string)

	go slowProcess(slow)
	go fastProcess(fast)

	// Wait for results with a timeout
	timeout := time.After(time.Second * 3)

	for i := 0; i < 2; i++ {
		select {
		case result := <-slow:
			fmt.Println(result)
		case result := <-fast:
			fmt.Println(result)
		case <-timeout:
			fmt.Println("Timed out waiting for a process")
			return
		}
	}
}
