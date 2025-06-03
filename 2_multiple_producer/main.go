/*
Exercise 2: Multiple Producers
Create a program with multiple producer goroutines sending data to a single channel,
and the main goroutine consuming the data.
*/
package main

import (
	"log"
	"time"
)

type messageData struct {
	producerOrigin string
	message        int
}

func main() {
	// Create single channel
	messageChannel := make(chan messageData)
	go producer(1, 5, messageChannel, "PRODUCERONE")
	go producer(6, 5, messageChannel, "PRODUCERTWO")

	for range 10 {
		messageData := <-messageChannel
		log.Println("CONSUMED : ", messageData)
	}

}

// Producer that sending data to a single channel
func producer(start, totalMessage int, messageChannel chan messageData, producerName string) {
	for i := range totalMessage {
		log.Println("SENDING FROM ", producerName)
		messageChannel <- messageData{
			producerOrigin: producerName,
			message:        start + i + 1,
		}
		time.Sleep(1 * time.Second)
	}

}
