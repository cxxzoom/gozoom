package main

import (
	"fmt"
	"time"
)

func main() {
	//go c1(1)
	//c1(2)

	s := []int{-5, -3, -4, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Printf("%d + %d = %d \n", x, y, x+y)
	cacheChan()
	time.Sleep(time.Second * 10)
}

func c1(a int) {
	for i := 0; i < 5; i++ {
		println(a)
		time.Sleep(100 * time.Millisecond)
	}
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func cacheChan() {
	c := make(chan int, 2)

	c <- 1
	c <- 2
	println(<-c)
	println(<-c)
}
