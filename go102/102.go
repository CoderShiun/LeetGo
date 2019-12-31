package go102

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var result [][]int

	bfs(root, &result, 0)

	return result
}

func bfs(root *TreeNode, result *[][]int, level int) {
	if root == nil {
		return
	}

	if level >= len(*result) {
		*result = append(*result, []int{})
	}

	(*result)[level] = append((*result)[level], root.Val)

	i := level + 1

	bfs(root.Left, result, i)
	bfs(root.Right, result, i)
}