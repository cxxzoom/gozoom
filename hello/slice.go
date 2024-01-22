package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	end := min(cap(arr), 5)
	slice := arr[0:end]
	println("len=", len(slice), "cap=", cap(arr))
	for k, v := range slice {
		println(k, " => ", v)
	}

	arr2 := make([]int, 3)
	for k, v := range arr2 {
		println(k, " => ", v)
	}

	s := []int{}
	if s == nil {
		println("empty")
	} else {
		println("\"not empty\"")
	}

	a := []int{1, 2, 3}
	ps(a)
	a = append(a, 4, 5, 6)
	ps(a)
	number := make([]int, 12)
	copy(number, a)
	ps(number)
	number = append(number, 123, 15)
	ps(number)

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i2, i3 := range pow {
		println("pow[", i2, "]=", i3, "\n")
	}

	m := map[int]string{
		1: "string",
	}

	for k, v := range m {
		println(k, "=>", v)
	}

	println(m[1])

	for i2, i3 := range "你好" {
		fmt.Printf("%d => %d\n", i2, i3)
	}
}

func ps(p []int) {
	//fmt.Printf("%d %d %d\n", len(p), cap(p), p)
}

func slice_sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}

	return sum
}
