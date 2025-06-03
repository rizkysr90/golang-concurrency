// Practice using select with multiple channels and a timeout.
package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan string, 0)
	timeoutCh := make(chan int, 0)
	// slowRequest := "SLOW_REQUEST"
	fastRequest := "FAST_REQUEST"
	go sentRequest(ch, fastRequest, timeoutCh)
	select {
	case value := <-ch:
		if value == "SLOW_REQUEST" {
			time.Sleep(5 * time.Second)
		} else {
			time.Sleep(1 * time.Second)
		}
	case second := <-timeoutCh:
		time.Sleep(time.Duration(second * int(time.Second)))
		log.Println("TIMEOUT")
		return
	}
}
func sentRequest(ch chan string, value string, timeoutCh chan int) {
	ch <- value
}
