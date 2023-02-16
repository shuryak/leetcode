/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    pQueue := []*TreeNode{p}
    qQueue := []*TreeNode{q}

    for len(pQueue) > 0 && len(qQueue) > 0 {
        pNode := pQueue[0]
        qNode := qQueue[0]
        pQueue = pQueue[1:]
        qQueue = qQueue[1:]

        if pNode == nil && qNode == nil {
            continue
        } else if pNode != nil && qNode != nil && pNode.Val == qNode.Val {
            pQueue = append(pQueue, pNode.Left, pNode.Right)
            qQueue = append(qQueue, qNode.Left, qNode.Right)
        } else {
            return false
        }
    }

    return true
}
