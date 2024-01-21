package main

func main() {
	for f := range fib(10) {
		println(f)
	}
}

func fib(length int) <-chan int {
	f := make(chan int, length)
	go func() {
		for i, j := 0, 1; i < length; i, j = j, i+j {
			f <- i
		}
		close(f)
	}()

	return f
}
