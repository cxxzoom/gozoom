package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	// test1()
	//defer printStack()
	//ffx(3)
	fmt.Println(testPanic())
	//fmt.Println(test111())
}

func test1() {
	switch s := drawCard(); s {
	case "Spades": // ...
	case "Hearts": // ...
	case "Diamonds": // ...
	case "Clubs": // ...
	default:
		panic(fmt.Sprintf("invalid suit %q", s)) // Joker?
	}
}

func drawCard() string {
	return "xxx"
}

func ffx(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	ffx(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func testPanic() (string, error) {
	type recoverP struct{}
	var err error
	defer func() {
		switch p := recover(); p {
		case nil:
		case recoverP{}:
			err = fmt.Errorf("some error")
		default:
			panic(p)
		}
	}()

	if "xxxx" != "xx" {
		panic(recoverP{})
		return "xxx", err
	}

	return "7777", err
}

func test111() (int, error) {
	return 7, fmt.Errorf("xxxxx")
}
