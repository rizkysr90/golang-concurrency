/*
Exercise 1: Simple Producer-Consumer
Create a program with one producer goroutine that sends numbers to a channel
and one consumer that reads and prints them.
*/
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	totalMessage := 5
	messageChannel := make(chan int)
	go func() {
		for i := range totalMessage {
			messageChannel <- i + 1
			time.Sleep(time.Millisecond * 100)
			log.Println("Sleep")
		}
		close(messageChannel)
	}()
	log.Println("Start")
	for message := range messageChannel {
		fmt.Println(message)
	}
	log.Println("Channel closed, finishing program")
}
