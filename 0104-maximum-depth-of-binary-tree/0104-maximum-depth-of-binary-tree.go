/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	searchQueue := []*TreeNode{root}

	depth := 0

	for len(searchQueue) > 0 {
		levelSize := len(searchQueue)

		for i := 0; i < levelSize; i++ {
			node := searchQueue[i]

			if node.Left != nil {
				searchQueue = append(searchQueue, node.Left)
			}
			if node.Right != nil {
				searchQueue = append(searchQueue, node.Right)
			}
		}

		depth++
		searchQueue = searchQueue[levelSize:]
	}

	return depth
}
