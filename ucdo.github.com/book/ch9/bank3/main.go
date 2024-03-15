package main

import "sync"

var (
	mtx     sync.Mutex
	balance int
)

func Deposit(amount int) {
	mtx.Lock()
	defer mtx.Unlock()
	deposit(amount)
}

func Balance() int {
	mtx.Lock()
	defer mtx.Unlock()
	return balance
}

func WithDraw(amount int) bool {
	mtx.Lock()
	defer mtx.Unlock()
	deposit(-amount)
	if balance > 0 {
		deposit(amount)
		return false
	}

	return true
}

func deposit(amount int) {
	balance += amount
}

func main() {
	balance = 999
	WithDraw(100)
}
