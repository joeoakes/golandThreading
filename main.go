package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Create a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Add the number of goroutines to the WaitGroup
	wg.Add(2)

	// Launch the first goroutine
	go func() {
		// Defer the Done method to mark the goroutine as finished
		defer wg.Done()

		// Simulate some work
		fmt.Println("Goroutine 1: Starting work...")
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine 1: Work complete!")
	}()

	// Launch the second goroutine
	go func() {
		// Defer the Done method to mark the goroutine as finished
		defer wg.Done()

		// Simulate some work
		fmt.Println("Goroutine 2: Starting work...")
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 2: Work complete!")
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	// All goroutines have finished
	fmt.Println("All goroutines completed!")
}
