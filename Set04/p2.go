package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
Problem 2: Implement a simple pub-sub system using channels.
*/

// buffered channel
func publish(i int, p chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	p <- i

}

func P2() {
	//fmt.Println("Hi p2")
	start := time.Now()
	pub := make(chan int, 100000000)
	var wg sync.WaitGroup

	for i := 1; i <= 100000000; i++ {
		wg.Add(1)
		go publish(i, pub, &wg)
	}

	go func() {
		wg.Wait()
		close(pub)
	}()

	//var swg sync.WaitGroup

	var counter int64 = 0
	// swg.Add(1)
	// go func(swg *sync.WaitGroup) {
	// 	defer swg.Done()
	// 	for range pub {
	// 		//fmt.Println(subscribe)
	// 		atomic.AddInt64(&counter, 1)
	// 	}
	// }(&swg)
	// swg.Wait()

	for range pub {
		//fmt.Println(subscribe)
		atomic.AddInt64(&counter, 1)
	}

	fmt.Println("elapsed time: ", time.Since(start))
	fmt.Println("full records: ", counter)
}

/////////// Mutex without buffer ////////////

// func publish(i int, pub chan int, wg *sync.WaitGroup, mu *sync.Mutex) {
// 	defer wg.Done()
// 	mu.Lock()
// 	pub <- i
// 	mu.Unlock()

// }

// func P2() {
// 	pub := make(chan int)
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex

// 	for i := 1; i <= 10; i++ {
// 		wg.Add(1)
// 		go publish(i, pub, &wg, &mu)
// 	}

// 	go func() {
// 		for p := range pub {
// 			fmt.Println(p)
// 		}
// 	}()

// 	wg.Wait()
// 	close(pub)

// }
