package main

import "sync"

func main() {
	//gen 0 - 9
	inputCh := genInput()

	// fan out squares calc
	squares1 := fanOut(inputCh)
	squares2 := fanOut(inputCh)

	// fan in results
	merged := fanIn(squares1, squares2)
	var sum int
	for v := range merged {
		sum += v
	}
	println(sum)
}

func genInput() <-chan int {
	in := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()
	return in
}

func fanOut(in <-chan int) <-chan int {
	squares := make(chan int)
	go func() {
		for v := range in {
			squares <- v * v
		}
		close(squares)
	}()
	return squares
}

func fanIn(in ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	wg.Add(len(in))

	merged := make(chan int)
	for _, ch := range in {
		ch := ch
		go func() {
			for v := range ch {
				merged <- v
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
