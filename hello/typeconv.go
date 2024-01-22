package main

import "strconv"

func main() {
	fl()
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
