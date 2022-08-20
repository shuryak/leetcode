/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    fast, slow := head, head
    left := []*ListNode{}
    
    if head.Next == nil {
        return true
    }
    
    var length int
    for head != nil {
        length++
        head = head.Next
    }
    
    for fast != nil && fast.Next != nil {
        left = append(left, slow)
        fast = fast.Next.Next
        slow = slow.Next
    }
    
    node := left[len(left)-1].Next
    for i := 0; node != nil; i++ {
        if length % 2 == 1 && i == 0 {
            node = node.Next
        }
        fmt.Println(i)
        if node.Val != left[len(left)-1-i].Val {
            return false
        }
        node = node.Next
    }
    return true
}
