package range_sum_of_bst

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	var sum int

	if root != nil {
		sumUp(root, L, R, &sum)
	}

	return sum
}

func sumUp(node *TreeNode, L, R int, sum *int) {
	if node.Val >= L && node.Val <= R {
		*sum += node.Val
	}

	if node.Left != nil {
		sumUp(node.Left, L, R, sum)
	}

	if node.Right != nil {
		sumUp(node.Right, L, R, sum)
	}
}