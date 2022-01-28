package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树前序遍历
func preorderTravelsal(root *TreeNode) []int {
	res := make([]int, 0)
	var preorder func(*TreeNode)
	//使用递归思想
	preorder = func(node *TreeNode)  {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return res
}
