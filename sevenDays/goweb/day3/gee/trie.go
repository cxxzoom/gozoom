// Package gee
package gee

// 设计一颗路由前缀树
import "strings"

// node 定义数的结构
type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}

// matchChild 这玩意儿是递归匹配的一环
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}

	return nil
}

// matchChildren 返回当前这一层的所有parts
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}

	return nodes
}

// insert 写入树
// 注意 child.insert()
// 这个就是写入到子树上面了
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
}

// search 不晓得是做啥子得，只看出来是递归
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasSuffix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
