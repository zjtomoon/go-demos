package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(maxDepth(root.Left),maxDepth(root.Right)) + 1
}


func Max(a,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	var dept int
	fmt.Println(dept)
}