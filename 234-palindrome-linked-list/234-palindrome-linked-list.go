/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    fast, slow := head, head
    
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    
    halfLength := getLength(slow)
    prev := slow
    cur := slow.Next
    
    for cur != nil {
        tmp := cur.Next
        cur.Next = prev
        prev = cur
        cur = tmp
    }
    
    left := head
    right := prev
    
    for i := 0; i < halfLength; i++ {
        if left.Val != right.Val {
            return false
        }
        left = left.Next
        right = right.Next
    }
    
    return true
}

func getLength(head *ListNode) int {
    var length int
    cur := head
    for cur != nil {
        cur = cur.Next
        length++
    }
    return length
}
