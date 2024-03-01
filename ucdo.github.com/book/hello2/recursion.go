package main

import "fmt"

type btree struct {
	value       int
	left, right *btree
}

func main() {
	s := []int{1, 2, 3, 3, 2, 1, -1, 2 - 1, 9, 9, 6, 4, 6}
	sort(s)
	fmt.Println(s)
}

func sort(values []int) {
	var root *btree
	for _, value := range values {
		root = genTree(root, value)
	}
	appends(root, values[:0])
}

func appends(root *btree, values []int) []int {
	if root != nil {
		values = appends(root.left, values)
		values = append(values, root.value)
		values = appends(root.right, values)
	}

	return values
}

func genTree(root *btree, value int) *btree {
	if root == nil {
		t := new(btree)
		t.value = value
		return t
	}

	// 比较大小看排那边
	if root.value > value {
		root.left = genTree(root.left, value)
	} else {
		root.right = genTree(root.right, value)
	}
	return root
}
