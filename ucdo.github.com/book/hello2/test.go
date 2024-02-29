package main

import "fmt"

func main() {
	a := ff()
	fmt.Println(a)
	a = ff1()
	fmt.Println(a)
	//fmt.Println(ff())
	//fmt.Println(ff1(ff()))
}

func ff() int {
	a := 1
	defer func() {
		a = 2
	}()

	return a
}

func ff1() (a int) {
	//a = 1
	defer func() {
		a = 2
	}()

	return a
}
