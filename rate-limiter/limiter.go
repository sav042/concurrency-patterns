package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// read values concurrently with 10 per second limit
func main() {
	count := 50
	results := make(chan int, count)
	lim := limiter()

	wg := sync.WaitGroup{}
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			results <- RPCWithLimiter(lim)
			wg.Done()
		}()
	}
	wg.Wait()
	close(results)

	for value := range results {
		fmt.Println(value)
	}
}

func RPCCall() int {
	return rand.Int()
}

func RPCWithLimiter(limiter <-chan struct{}) int {
	<-limiter
	return RPCCall()
}

func limiter() chan struct{} {
	lim := make(chan struct{}, 10)
	go func() {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-t.C:
				for i := 0; i < 10; i++ {
					lim <- struct{}{}
				}
			}
		}
	}()
	return lim
}
