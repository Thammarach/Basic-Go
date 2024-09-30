package main

import (
	"fmt"
	"time"
)

func main() {
	jobCh := make(chan int, 10)
	resultCh := make(chan int, 10)

	// Use a standard for loop to send jobs to the jobCh channel
	for i := 0; i < 10; i++ {
		jobCh <- i + 1
	}
	close(jobCh)

	// Start two goroutines for processing jobs
	for i := 0; i < 2; i++ {
		go double(jobCh, resultCh)
	}

	// Collect and print results
	for i := 0; i < 10; i++ {
		result := <-resultCh
		fmt.Println(result)
	}
}

func double(jobCh <-chan int, resultCh chan<- int) {
	for j := range jobCh {
		time.Sleep(1 * time.Second)
		resultCh <- j * 2
	}
}
