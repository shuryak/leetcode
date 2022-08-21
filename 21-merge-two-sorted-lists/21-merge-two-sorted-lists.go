/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var lead, driven *ListNode
    
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }
    
    if list1.Val <= list2.Val {
        lead, driven = list1, list2   
    } else {
        driven, lead = list1, list2   
    }
    
    head := lead
    
    for lead != nil && lead.Next != nil {
        if lead.Val <= driven.Val && lead.Next.Val >= driven.Val {
            tmp := lead.Next
            lead.Next = driven
            lead = driven
            driven = tmp
        } else {
            lead = lead.Next
        }
    }
    
    lead.Next = driven   
    
    return head
}
