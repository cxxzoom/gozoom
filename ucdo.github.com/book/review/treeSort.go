package main

import "fmt"

type btree struct {
	value       int
	left, right *btree
}

func main() {
	s := []int{1, 9, 8, 5, 0, -1, -6, 8, 9}
	sort(s)
	fmt.Println(s)
}

func sort(values []int) {
	var tree *btree
	for _, value := range values {
		tree = genTree(tree, value)
	}
	// 前序遍历以排序
	appendValue(tree, values[:0])
}

func appendValue(tree *btree, values []int) []int {
	if tree != nil {
		values = appendValue(tree.left, values)
		values = append(values, tree.value)
		values = appendValue(tree.right, values)
	}

	return values
}

func genTree(tree *btree, value int) *btree {
	// gen a node
	if tree == nil {
		t := new(btree)
		t.value = value
		return t
	}
	// 比较当前值和当前节点的值
	if value < tree.value {
		tree.left = genTree(tree.left, value)
	} else {
		tree.right = genTree(tree.right, value)
	}
	return tree
}
