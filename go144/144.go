package go144

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var result []int
	appendNode(root, &result)
	return result
}

func appendNode(node *TreeNode, arr *[]int) {
	if node == nil {
		return
	}

	*arr = append(*arr, node.Val)

	if node.Left != nil {
		appendNode(node.Left, arr)
	}

	if node.Right != nil {
		appendNode(node.Right, arr)
	}
}