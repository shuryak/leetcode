/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    avgs := []float64{float64(root.Val)}

    if root.Left == nil && root.Right == nil {
        return avgs
    }
    
    
    searchQueue := []*TreeNode{root.Left, root.Right}

    for len(searchQueue) > 0 {
        levelSize := len(searchQueue)
        nodesOnLevel := 0

        sum := 0.0

        for i := 0; i < levelSize; i++ {
            node := searchQueue[i]
            if node == nil {
                continue
            }

            nodesOnLevel++
            sum += float64(node.Val)

            if node.Left != nil {
                searchQueue = append(searchQueue, node.Left)
            }
            if node.Right != nil {
                searchQueue = append(searchQueue, node.Right)
            }
        }

        avgs = append(avgs, sum / float64(nodesOnLevel))
        searchQueue = searchQueue[levelSize:]
    }

    return avgs
}
