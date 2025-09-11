package main

import (
	"fmt"
	"time"
)

/*
  Bonus: Implement a semaphore using channels.
*/

func Bonus() {

	sem := make(chan struct{}, 1)

	for i := 1; i <= 5; i++ {
		go func(i int) {
			sem <- struct{}{}

			fmt.Printf("Goroutine %d running\n ", i)
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("Goroutine %d is releasing\n ", i)
			<-sem
		}(i)
	}
	// Wait enough time for all goroutines to finish
	time.Sleep(6 * time.Second)
	fmt.Println("All goroutines done")
}
