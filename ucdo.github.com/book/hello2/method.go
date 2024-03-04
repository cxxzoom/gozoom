package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
}

func Distance(p, q Point) float64 {
	return math.Hypot(p.x-q.x, p.y-q.y)
}

// 这里的p称之为方法接收器，可以在上面搞很多方法
// 在其他语言里，this,self也被称之为方法接收器
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.x-q.x, p.y-q.y)
}

type p *int

// func (p) p1() {} // 如果原始类型是指针，则没办法当作接收器
