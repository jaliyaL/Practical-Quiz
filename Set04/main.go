package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const total = 100_000_000

func publish1(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		ch <- i
	}
}

func main() {
	fmt.Println("Benchmarking...")

	// -------------------------
	// 1. Reader in main goroutine
	// -------------------------
	{
		start := time.Now()
		ch := make(chan int, 100000)
		var wg sync.WaitGroup
		wg.Add(1)
		go publish1(total, ch, &wg)

		// Close channel after publishing
		go func() {
			wg.Wait()
			close(ch)
		}()

		var counter int64 = 0
		for range ch {
			atomic.AddInt64(&counter, 1)
		}

		elapsed := time.Since(start)
		fmt.Printf("Main goroutine reader: elapsed=%v, counter=%d\n", elapsed, counter)
	}

	// -------------------------
	// 2. Reader in separate goroutine
	// -------------------------
	{
		start := time.Now()
		ch := make(chan int, 100000)
		var wg sync.WaitGroup
		wg.Add(1)
		go publish1(total, ch, &wg)

		// Close channel after publishing
		go func() {
			wg.Wait()
			close(ch)
		}()

		var counter int64 = 0
		var swg sync.WaitGroup
		swg.Add(1)
		go func() {
			defer swg.Done()
			for range ch {
				atomic.AddInt64(&counter, 1)
			}
		}()
		swg.Wait()

		elapsed := time.Since(start)
		fmt.Printf("Reader in goroutine: elapsed=%v, counter=%d\n", elapsed, counter)
	}
}
