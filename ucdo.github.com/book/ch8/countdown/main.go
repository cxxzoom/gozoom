package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// countdown1()
	abort := make(chan struct{})
	// countdown2()
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	tick := time.Tick(1 * time.Second)
	for i := 10; i > 0; i-- {
		select {
		case <-tick:
		case <-abort:
			fmt.Println("abort!!")
			return
		}
		launch()
	}
	// test1()
}

func launch() {
	fmt.Println("launch!!")
}

func countdown1() {
	fmt.Println("countdown 1")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func countdown2() {
	fmt.Println("countdown 2")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	// launch()
}

func test1() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
