package main

import (
	"fmt"
	"strconv"
)

func main() {
	fl()
	s_conv()
	interfaceC()
}

func fl() {
	a := 1.00100
	b := int(a)
	println(a, b)

	c := "sadjfo"
	d, err := strconv.Atoi(c)
	if err != nil {
		println("is not a numeric")
	} else {
		println(d)
	}

}

func s_conv() {
	str := "3.14"
	s, err := strconv.ParseFloat(str, 64)
	if err != nil {
		println("is not a float string")
	} else {
		println(s)
	}

	f := 3.14
	q := strconv.FormatFloat(f, 'f', 10, 64)
	println(q)
}

func interfaceC() {
	var i interface{} = "Hello, World"
	str, ok := i.(int)
	if ok {
		fmt.Printf("'%s' is a string\n", str)
	} else {
		fmt.Println("conversion failed")
	}
}
