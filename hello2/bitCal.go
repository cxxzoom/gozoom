package main

import "fmt"

func main() {
	var x uint8 = 1 << 1
	fmt.Printf("%8b\n", x)
	var x1 uint8 = 1 << 5
	fmt.Printf("%8b\n", x1)
	fmt.Printf("%8b\n", x|x1)

	o := 0666
	fmt.Printf("%d %o %#o\n", o, o, o)

	xx := 0xabcde
	fmt.Printf("%d %x %#x\n", xx, xx, xx)

	b := '国'
	fmt.Printf("%d %c %q", b, b, b)
}
