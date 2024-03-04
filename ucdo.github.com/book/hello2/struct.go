package main

import (
	"fmt"
	"net/url"
)

type employee struct {
	ID       int
	Name     string
	Salary   int
	Position string
}

type point struct {
	x, y int
}

type circle struct {
	point  point
	Radius int
}

type wheel struct {
	circle circle
	spokes int
}

type point2 struct {
	x, Y int
}

type circle2 struct {
	point2
	Radius int
}

type wheel2 struct {
	circle2
	Spokes int
}

var e employee

func main() {
	println(e.ID)
	f := &e
	f.ID = 7777
	id := &f.ID
	*id = 9999
	position := &f.Position
	*position = "sdfksjdjfk"
	println(e.ID)
	println(e.Position)
	fmt.Println(getId(&e).ID)
	getId(&e).ID = 9
	fmt.Println(e.ID)

	values := []int{15, 4, 3, 2, 1, 8, 2, 234, 5, 567}
	sort2(values)
	fmt.Println(values)

	fmt.Println("++++++++++++++++++++++++++++")
	fmt.Println(e)
	editS(&e)
	fmt.Println(e)

	// 初始化并获取地址
	_ = &employee{ID: 111}
	pp1 := new(employee)
	*pp1 = employee{ID: 1211}
	fmt.Println("++++++++++++++++++++++++++++")
	mm := make(map[employee]int)
	key := employee{
		ID:       0,
		Name:     "",
		Salary:   0,
		Position: "",
	}
	mm[key]++
	for k, v := range mm {
		fmt.Println(k, v)
	}
	key.ID = 11111

	for k, v := range mm {
		fmt.Println(k.ID, v)
	}

	w := wheel{
		circle: circle{
			point: point{
				x: 1,
				y: 11,
			},
			Radius: 2,
		},
		spokes: 8,
	}

	fmt.Printf("%#v", w)
	var ffff wheel2
	ffff.x = 1
}

func getId(id *employee) *employee {
	id.ID = 0001
	return id
}

func editS(s *employee) {
	s.Name = "xxxxxxx"
}

type tree2 struct {
	value       int
	left, right *tree2
}

// 思路： 先把值加入到树里面，然后进行排序

func sort2(values []int) {
	var root *tree2
	for _, value := range values {
		root = add2(value, root)
	}
	appendValues(values[:0], root)
}

func add2(value int, tree *tree2) *tree2 {
	if tree == nil {
		tmp := new(tree2)
		tmp.value = value
		return tmp
	}
	if value < tree.value {
		tree.left = add2(value, tree.left)
	} else {
		tree.right = add2(value, tree.right)
	}
	return tree
}

func appendValues(value []int, tree *tree2) []int {
	if tree != nil {
		value = appendValues(value, tree.left)
		value = append(value, tree.value)
		value = appendValues(value, tree.right)
	}

	return value
}

// An intList  is a linked list of integers
// A nil *intList represents empty list
type intList struct {
	Value int
	Tail  *intList
}

func (list *intList) sum() int {
	if list == nil {
		return 0
	}

	return list.Value + list.Tail.Value
}

func urlValues() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang"))

	fmt.Println(m.Get("lang"))

	fmt.Println(m.Get("lang"))

	fmt.Println(m.Get("item"))
}
