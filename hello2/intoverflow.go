package main

import "fmt"

func main() {
	var u uint8 = 255

	var i int8 = 127

	fmt.Println(u, u*u, u+u)

	fmt.Println(i, i+1, i*i, i-127-127)
}
