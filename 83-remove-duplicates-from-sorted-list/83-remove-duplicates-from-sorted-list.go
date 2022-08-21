/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    
    prev := head
    cur := head.Next
    
    for cur != nil {
        if prev.Val == cur.Val {
            if prev == head {
                head = cur
                prev = cur
                cur = cur.Next
            } else {
                prev.Next = cur.Next
                cur = cur.Next
            }
        } else {
            prev = cur
            cur = cur.Next
        }
    }
    
    return head
}