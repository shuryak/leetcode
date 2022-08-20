/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    var prev *ListNode
    cur := head
    
    for cur != nil {
        if cur.Val != val {
            prev = cur
            cur = cur.Next
            continue
        }
        
        if prev == nil { // у элемента нет предыдущего => это первый элемент
            head = head.Next
        } else {
            prev.Next = cur.Next // выкинули элемент cur из цепочки
        }
        
        cur = cur.Next
    }
    
    return head
}
