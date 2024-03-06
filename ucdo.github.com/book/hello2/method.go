package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

type Point struct {
	x, y float64
	X, Y float64
}

type ColorPoint struct {
	Point
	Color color.RGBA
}

type ColorPoint2 struct {
	*Point
	//ColorPoint
	Color color.RGBA
}

func main() {
	//p := Point{1, 2}
	//q := Point{4, 6}
	//fmt.Println(Distance(p, q))
	//fmt.Println(p.Distance(q))
	// 这里通过var定义，只能只能显示赋值
	//var cp ColorPoint
	//cp.x = 1
	//cp.y = 2
	//// 下面这种定义方式：必须显示指定
	//cp2 := ColorPoint{
	//	Point: Point{
	//		x: 1,
	//		y: 2,
	//	},
	//}
	//fmt.Println(cp2)
	//red := color.RGBA{255, 0, 0, 255}
	//blue := color.RGBA{0, 0, 255, 255}
	//cp.Color = red
	//cp.Color = blue

	//var cp1 ColorPoint2
	//cp1.x = 1
	//cp1 := ColorPoint2{
	//	Point: &Point{
	//		x: 0,
	//		y: 0,
	//	},
	//}
	//
	//cp3 := ColorPoint2{Point: &Point{x: 1, y: 2}, Color: color.RGBA{B: 255, A: 255}}
	//fmt.Println(cp1.Distance(*cp1.Point))
	//fmt.Println(cp3.Distance(*cp3.Point))
	//a := ColorPoint2{&Point{x: 1, y: 2}, color.RGBA{0, 0, 0, 0}}
	//a.Distance(*a.Point)

	//fmt.Println(cp.Distance(cp)) // Cannot use 'cp' (type ColorPoint) as the type Point
	p := Point{1, 1, 1, 1}
	path := Path{{1, 1, 2, 2}, {0, 0, 0, 0}}
	path.xxxx(p, true)
	fmt.Println(path)
}

func Distance(p, q Point) float64 {
	return math.Hypot(p.x-q.x, p.y-q.y)
}

// 这里的p称之为方法接收器，可以在上面搞很多方法
// 在其他语言里，this,self也被称之为方法接收器
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.x-q.x, p.y-q.y)
}

func (p *Point) Salary(s int) {

}

type p *int

// func (p) p1() {} // 如果原始类型是指针，则没办法当作接收器

// anonymous struct
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}

// 下面的代码就不算很难，相当于是
func (p Point) Add(q Point) Point { return Point{x: p.X + q.X, y: p.Y - q.Y} }
func (p Point) Sub(q Point) Point { return Point{x: p.X - q.X, y: p.Y - q.Y} }

type Path []Point

func (path Path) xxxx(offset Point, isAdd bool) {
	var op func(p, q Point) Point
	if isAdd {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}
