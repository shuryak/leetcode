/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type NodePath struct {
	*TreeNode
	Sum int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	searchStack := []*NodePath{{root, root.Val}}

	for len(searchStack) > 0 {
		node := searchStack[len(searchStack)-1]
		searchStack = searchStack[:len(searchStack)-1]

		isLeaf := node.Left == nil && node.Right == nil
		if isLeaf && node.Sum == targetSum {
			return true
		}

		if node.Right != nil {
			searchStack = append(
				searchStack,
				&NodePath{node.Right, node.Sum + node.Right.Val},
			)
		}
		if node.Left != nil {
			searchStack = append(
				searchStack,
				&NodePath{node.Left, node.Sum + node.Left.Val},
			)
		}
	}

	return false
}
