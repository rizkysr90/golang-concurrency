/*
Exercise 3: Channel Coordination
Use channels for coordination between goroutines instead of using sync.WaitGroup.
*/
package main

import (
	"log"
)

func main() {
	courierChannel := make(chan string)
	cookChannel := make(chan string)
	done := make(chan bool) // Done signal channel

	go courierWorker(courierChannel)
	go cookWorker(courierChannel, cookChannel)
	go waiterWorker(cookChannel, done)

	<-done // Wait for completion signal
	log.Println("All work completed!")
}

func courierWorker(courierChannel chan string) {
	cakeIngredientsToShop := []string{
		"eggs",
		"flour",
		"butter",
		"chocolate",
		"milk",
	}
	for _, cakeIngredient := range cakeIngredientsToShop {
		log.Println("Courier worker : sent ", cakeIngredient)
		courierChannel <- cakeIngredient
	}
	close(courierChannel)
}
func cookWorker(courierChannel, cookChannel chan string) {
	for value := range courierChannel {
		log.Println("Cook worker : received ", value)
		cookChannel <- value
	}
	close(cookChannel)
}

func waiterWorker(cookChannel chan string, done chan bool) {
	for value := range cookChannel {
		log.Println("Waiter worker : received ", value)
	}
	done <- true // Signal completion
}
