/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil && root2 == nil {
        return nil
    }
    if root1 == nil {
        return root2
    }
    if root2 == nil {
        return root1
    }

    stack1 := []*TreeNode{root1}
    stack2 := []*TreeNode{root2}

    for len(stack1) > 0 {
        node1 := stack1[len(stack1)-1]
        node2 := stack2[len(stack2)-1]
        stack1 = stack1[:len(stack1)-1]
        stack2 = stack2[:len(stack2)-1]

        node1.Val += node2.Val

        if node1.Left == nil && node2.Left != nil {
            node1.Left = node2.Left
        } else if node1.Left != nil && node2.Left != nil {
            stack1 = append(stack1, node1.Left)
            stack2 = append(stack2, node2.Left)
        }

        if node1.Right == nil && node2.Right != nil {
            node1.Right = node2.Right
        } else if node1.Right != nil && node2.Right != nil {
            stack1 = append(stack1, node1.Right)
            stack2 = append(stack2, node2.Right)
        }
    }

    return root1
}
