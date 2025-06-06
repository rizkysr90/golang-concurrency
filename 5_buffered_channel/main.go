// Use a buffered channel to limit concurrency without sync.Semaphore.

package main

import (
	"log"
	"time"
)

func main() {
	// Buffered channel acts as semaphore - limits to 5 concurrent workers
	semaphore := make(chan struct{}, 5)
	// Channel to collect results
	results := make(chan string)

	foods := []string{
		"pecel ayam", "pecel lele", "rendang", "sate",
		"soto", "bakso", "mie",
	}

	// Launch all goroutines
	for _, food := range foods {
		go func(f string) {
			// Acquire semaphore (blocks if 5 goroutines already running)
			semaphore <- struct{}{}

			// Do the actual work
			processFood(f)

			// Release semaphore and send result
			<-semaphore
			results <- f
		}(food)
	}
	// Collect all results
	timeout := time.After(60 * time.Second)
	for i := 0; i < len(foods); i++ {
		select {
		case food := <-results:
			log.Println("finish processing food:", food)
		case <-timeout:
			log.Println("timeout")
			return
		}
	}
}

func processFood(food string) {
	log.Println("processing food:", food)
	time.Sleep(2 * time.Second)
}
