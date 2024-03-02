package main

import (
	"fmt"
	"sort"
)

// https://golang-china.github.io/gopl-zh/ch5/ch5-06.html
// 有环则不能进行拓扑排序
// 拓扑排序
// 从概念上说，前置条件可以构成有向图
// 从该点出发，不会再回到该点，就是说明是无环
func main() {
	var collection = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},
		"compilers": {
			"data structures",
			"formal languages",
			"computer organization",
		},
		"data structures":       {"discrete math"},
		"databases":             {"data structures"},
		"discrete math":         {"intro to programming"},
		"formal languages":      {"discrete math"},
		"networks":              {"operating systems"},
		"operating systems":     {"data structures", "computer organization"},
		"programming languages": {"data structures", "computer organization"},
	}

	for k, v := range toposort(collection) {
		fmt.Println(k, v)
	}
}

func toposort(val map[string][]string) []string {
	var order []string

	var keys []string
	for key := range val {
		keys = append(keys, key)
	}
	//fmt.Println(order)
	ex := make(map[string]bool) // 判断值值是否存在
	var visitAll func(items []string)
	visitAll = func(items []string) { // 这里传的keys
		for _, item := range items {
			if !ex[item] {
				ex[item] = true
				visitAll(val[item])
				order = append(order, item)
			}
		}
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}
