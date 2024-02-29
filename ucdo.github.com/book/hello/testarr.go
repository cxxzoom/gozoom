package main

import "fmt"

func main() {
	//a := [11]int{1, 2, 3}
	// t(a) 错误的，因为只要长度为10的数组
	c := closure()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := closure()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	fmt.Println("test get area")
	d := Circle{}
	d.radius = 10
	fmt.Println(d.getArea())

	fmt.Println("test t func")

	t()

	fmt.Println("test t tslice")
	tslice()

	fmt.Println("test t tt")
	tt()
}

func t() {
	a := []int{1: 1, 3: 9}
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}

func closure() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func tslice() {
	a := [][]int{
		{1, 2, 3, 4, 5},
		{2, 3, 4, 5, 6},
	}
	for key, val := range a {
		println(key)
		for k, v := range val {
			println(k, " => ", v)
		}
	}
}

const Max = 3

func tt() {
	a := [Max]int{1, 2, 3}
	var ptr [Max]*int
	for i := 0; i < len(a); i++ {
		ptr[i] = &a[i]
		println(a[i])
	}

	for k, v := range ptr {
		println(k, "=>", *v)
	}

}
