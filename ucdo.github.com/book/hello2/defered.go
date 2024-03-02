package main

import (
	"log"
	"time"
)

func main() {
	bigSlowOperation()
}

func bigSlowOperation() {
	defer trace("operations")()
	time.Sleep(time.Second * 10)
}

func trace(msg string) func() {
	timeStart := time.Now()
	log.Printf("%s is start...\n", msg)
	return func() {
		log.Printf("%s is ending, cost time %s", msg, time.Since(timeStart))
	}
}

func double2(a int) (res int) {
	defer func(){
		res += a
	}()
	return a + a
}
