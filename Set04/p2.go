package main

import (
	"fmt"
	"sync"
)

/*
Problem 2: Implement a simple pub-sub system using channels.
*/

// buffered channel
// func publish(i int, p chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	p <- i

// }

// func P2() {
// 	//fmt.Println("Hi p2")
// 	pub := make(chan int, 10)
// 	var wg sync.WaitGroup

// 	for i := 1; i <= 10; i++ {
// 		wg.Add(1)
// 		go publish(i, pub, &wg)
// 	}

// 	wg.Wait()
// 	close(pub)

// 	for subscribe := range pub {
// 		fmt.Println("values", subscribe)
// 	}
// }

/////////// Mutex without buffer ////////////

func publish(i int, pub chan int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	pub <- i
	mu.Unlock()

}

func P2() {
	pub := make(chan int)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go publish(i, pub, &wg, &mu)
	}

	go func() {
		for p := range pub {
			fmt.Println(p)
		}
	}()

	wg.Wait()
	close(pub)

}
