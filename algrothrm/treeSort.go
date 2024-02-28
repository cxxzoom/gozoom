package main

import "fmt"

// 这是个使用二叉树进行插入排序的算法
// 1. 思路： 先把值构造成一个二叉树
// 2. 再对这个二叉树进行前序遍历
// 3. 构造二叉树的思路： 先创建一个结构体，用这个结构体来充当一棵树，类似 C的链表，左右节点分别为指向自身的指针
// 4. 然后对要排序的值进行遍历：这里创建一个新的函数，因为会用用到递归
// 5. 构造二叉树的时候，如果根节点是空的，就new(T) 然后把赋值
// 6. 比较根节点和当前值的大小，如果当前值小于当前节点的值，那么当前节点的左节点就为这个值 （这里注意要用递归，因为左节点是空的，要创建）
// 7. 右节点如法炮制
// 8. 然后进行树的前序遍历了
// 9. 节点不为空，则进行左边子树的递归，然后赋值（这里是中间节点），然后再进行右边子树的递归
// 10. 因为是slice，引用了底层数组，所以就直接对原数组进行了操作

type tree struct {
	value       int
	left, right *tree
}

func main() {
	values := []int{9, 2, 1, 5, 6, 3, 8, 2, 0, 9, -1, -3, -5}
	sort(values)
	fmt.Println(values)
}

func sort(values []int) {
	var root *tree
	for _, v := range values {
		// 这里把值放到树里面
		root = genTree(root, v)
	}
	// 进行前序遍历来排序
	// 这里的values[:0] 意思是清空values这个切片
	appendValue(values[:0], root)
}

func appendValue(values []int, root *tree) []int {
	// 如果节点不为空，则进行递归进行值的写入
	if root != nil {
		values = appendValue(values, root.left)
		values = append(values, root.value)
		values = appendValue(values, root.right)
	}
	// 为空就直接返回
	return values
}

func genTree(root *tree, v int) *tree {
	// 如果这棵树是空的
	if root == nil {
		t := new(tree)
		t.value = v
		return t
	}
	// 比较根节点的value待写入的值
	// 小于根节点，写在左边；否则写在右边
	if root.value > v {
		root.left = genTree(root.left, v)
	} else {
		root.right = genTree(root.right, v)
	}

	return root
}
