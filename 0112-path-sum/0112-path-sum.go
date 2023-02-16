/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil{
        return false
    }

    sums := make(map[*TreeNode]int)

    stack := []*TreeNode{root}

    for len(stack) > 0 {
        for i := 0; i < len(stack); i++ {
            node := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            if node == nil {
                continue
            }

            isLeaf := node.Left == nil && node.Right == nil

            if isLeaf && sums[node] + node.Val == targetSum {
                return true
            }

            if !isLeaf {
                sums[node.Left] = sums[node] + node.Val
                sums[node.Right] = sums[node] + node.Val
            }

            stack = append(stack, node.Right, node.Left)
        }
    }

    return false
}
