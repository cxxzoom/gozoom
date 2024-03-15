package main

var (
	sema    = make(chan struct{}, 1)
	balance int
)

/*
这里使用了单个buffer的控制
即每个用的时候都要往channel里面写数据
但是如果有其他人在用，就会被阻塞
依次来保证了并发安全
*/
func Deposit(amount int) {
	sema <- struct{}{}
	balance += amount
	<-sema
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
