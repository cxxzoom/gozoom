package main

import "fmt"

type employee struct {
	ID       int
	Name     string
	Salary   int
	Position string
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
}

func getId(id *employee) *employee {
	id.ID = 0001
	return id
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
