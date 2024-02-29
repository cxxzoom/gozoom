package main

import "fmt"

// Ab 定义结构体
type Ab interface {
	write()
}

type AC interface {
	print()
}

type B struct {
}

func (b *B) write() {
	println("I`m part of interface Ab")
}

func (b *B) print() {

}

type phone interface {
	call()
	send()
}

type nPhone struct {
}

type iPhone struct {
}

func (n nPhone) call() {
	println("i am nigger")
}

func (i iPhone) call() {
	println("i m iphone")
}

type area interface {
	area()
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.width
}

type CCircle struct {
	radius float64
}

const PI = 3.14

func (c CCircle) area() float64 {
	return PI * c.radius * c.radius
}

// 定义一个错误的接口

type errDeal struct {
	a int
	b int
}

func (e errDeal) err() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, e.a)
}

func (e errDeal) div() (int, string) {
	a, b := e.a, e.b
	if b == 0 {
		return 0, e.err()
	}

	return a / b, ""
}

func main() {
	n := nPhone{}
	n.call()

	m := new(nPhone)
	m.call()

	r := Rectangle{width: 1, height: 1}
	fmt.Printf("area is %f\n", r.area())

	f := CCircle{radius: 1.3}
	fmt.Printf("area is %f\n", f.area())

	t := errDeal{1, 0}
	if _, err := t.div(); err != "" {
		println(err)
	}

	ff := errDeal{2, 2}
	if tt, err := ff.div(); err == "" {
		println(tt)
	}

}
