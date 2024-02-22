package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	t1()
	t2()
	t3()
	t4()
}

func t1() {
	println("t1....")
	var a [3]int
	for i, i2 := range a {
		println(i, " => ", i2)
	}
}

func t2() {
	println("t2....")
	type Currency int
	const (
		U Currency = iota
		E
		G
		R
	)

	symbol := [...]string{U: "U", E: "E", G: "G", R: "R"}
	fmt.Println(R, symbol[R])
}

func t3() {
	fmt.Println("t3...test two slice is equ")
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Println(c1)
	fmt.Printf("%x\n%x\n%t\n%T", c1, c2, c1 == c2, c1)
}

func t4() {
	fmt.Println("t4....")
	arr := [4]int{1, 2, 3, 4}
	for _, v := range arr {
		fmt.Println(v)
	}

	clearArray(&arr)
	fmt.Println("clear array...")

	for _, v := range arr {
		fmt.Println(v)
	}

	arr = [4]int{}
}

func clearArray(ptr *[4]int) {
	for i := range ptr {
		ptr[i] = 0
	}

	for i := range ptr {
		ptr[i] = 0
	}

	*ptr = [4]int{}
}
