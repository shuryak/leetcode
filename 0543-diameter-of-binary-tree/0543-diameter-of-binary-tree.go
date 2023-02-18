/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }

    depths := make(map[*TreeNode]int)
    stack := []*TreeNode{root}

    maxDiameter := 0

    for len(stack) > 0 {
        node := stack[len(stack)-1]

        if node.Left != nil && depths[node.Left] == 0 { // Не посещали левую ноду
            stack = append(stack, node.Left)
        } else if node.Right != nil && depths[node.Right] == 0 { // Не посещали правую ноду
            stack = append(stack, node.Right)
        } else {
            stack = stack[:len(stack)-1]

            curDiameter := depths[node.Left] + depths[node.Right]
            if curDiameter > maxDiameter {
                maxDiameter = curDiameter
            }

            if depths[node.Left] >= depths[node.Right] {
                depths[node] = depths[node.Left] + 1
            } else {
                depths[node] = depths[node.Right] + 1
            }
        }
    }

    return maxDiameter
}
