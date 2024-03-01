package main

import (
	"fmt"
)

func main() {
	//f := negative
	//f(1)
	//fmt.Printf("%T\n", f)
	//f = square
	////f = product error
	//fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "hello world!"))

	f := squares
	// () 函数调用符号
	fmt.Println(f()()) // "1"
	fmt.Println(f()()) // "4"
	fmt.Println(f()()) // "9"
	fmt.Println(f()()) // "16"
}

func negative(a int) int {
	return -a
}

func square(n int) int     { return n * n }
func product(m, n int) int { return m * n }

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
