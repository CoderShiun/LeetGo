package go145

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var result []int

	appendNode(root, &result)

	return result
}

func appendNode (root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		appendNode (root.Left, result)
	}

	if root.Right != nil {
		appendNode (root.Right, result)
	}

	*result = append(*result, root.Val)
}