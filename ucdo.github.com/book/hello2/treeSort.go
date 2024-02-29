package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

// 二叉树插入排序
func main() {
	values := []int{15, 4, 3, 2, 1, 8, 2, 234, 5, 567}
	Sort(values)
	fmt.Println(values)
}

func Sort(value []int) {
	var root *tree
	for _, v := range value {
		root = add(root, v)
	}
	appendValue(value[:0], root)
}

func appendValue(val []int, t *tree) []int {
	if t != nil {
		val = appendValue(val, t.left)
		val = append(val, t.value)
		val = appendValue(val, t.right)
	}
	return val
}

// 构造一棵树：
func add(val *tree, value int) *tree {
	if val == nil {
		tmp := new(tree) // 创建一棵新树
		tmp.value = value
		return tmp
	}
	if val.value < value {
		val.right = add(val.right, value)
	} else {
		val.left = add(val.left, value)
	}
	return val
}
