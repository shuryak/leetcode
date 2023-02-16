/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    searchQueue := []*TreeNode{root}

    depth := 1

    for len(searchQueue) > 0 {
        levelSize := len(searchQueue) // depth нужно увеличивать только когда пробежимся по всему "горизонтальному уровню"

        for i := 0; i < levelSize; i++ {
            node := searchQueue[0]
            searchQueue = searchQueue[1:]

            if node != nil {
                searchQueue = append(searchQueue, node.Left, node.Right)
            } else {
                continue
            }

            if node.Left == nil && node.Right == nil {
                return depth
            }
        }
        
        depth++
    }

    return depth
}
