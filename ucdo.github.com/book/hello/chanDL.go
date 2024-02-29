package main

// 演示： deadlock in chan with buffer
// fatal error: all goroutines are asleep - deadlock!
//
// goroutine 1 [chan send]:
// main.main()
//
//	F:/go/src/hello/chanDL.go:7 +0x47
//
// exit status 2

// 为什么？ 因为缓冲区满了，但是这时没人来读，就导致阻塞，就死锁了
func main() {
	c := make(chan int, 1)
	c <- 1
	c <- 2
	println(<-c)
	println(<-c)
}
