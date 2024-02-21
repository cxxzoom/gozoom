package main

import "fmt"

func main() {
	s := "国家"
	s1 := s[1:4]

	fmt.Println(len(s), s1)

	t := "123"
	xx := t
	t = "0000" + t
	t += "0000"

	println(t)
	println(xx)
}
