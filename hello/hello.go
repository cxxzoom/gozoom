package main

import (
	"fmt"
	"unsafe"
)

const (
	a = "string"
	b = len(a)
	c = unsafe.Sizeof(a)
)

const (
	i = 1 << iota
	j = 3 << iota
	k // 实际上是 3 << 2
	f // 实际上是 3 << 3
)

func main() {
	var a *int
	var b **int
	var c ***int
	var f *int
	d := 1
	a = &d
	b = &a
	c = &b
	fmt.Printf("d = %d", ***c) //caddress = c00000a0c8 aaddress = c00004e020 value = 1
	if f == nil {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}

	g := 1
	h := 2
	swap2(&g, &h)
	fmt.Println(g, h)

}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func swap(a, b string) (string, string) {
	a, b = b, a

	return a, b
}

func swap2(a *int, b *int) {
	fmt.Printf("%x %x\n", a, b)
	tmp := a
	a = b
	b = tmp
}
