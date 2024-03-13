package main

import "fmt"

// send msg to channel
func counter(in chan<- int) {
	for i := 0; i < 100; i++ {
		in <- i
	}
	close(in)
}

// get msg on channel and then send msg to channel
func squarer(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

// get msg on channel and print
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	s := make(chan int)
	n := make(chan int)
	go counter(s)
	go squarer(s, n)
	printer(n)
}
