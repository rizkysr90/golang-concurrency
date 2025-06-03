/*
Exercise 3: Channel Coordination
Use channels for coordination between goroutines instead of using sync.WaitGroup.
*/
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, done chan<- bool) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Millisecond * 300)
	}
	// Signal that this worker is done
	done <- true
}

func main() {
	jobs := make(chan int, 10)
	done := make(chan bool)

	// Create 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, done)
	}

	// Send 10 jobs
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for all three workers to finish
	for i := 0; i < 3; i++ {
		<-done
	}
	fmt.Println("All workers completed their jobs")
}
