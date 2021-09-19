package linked_list

func RemoveElements(head *ListNode, val int) *ListNode {
    dummy := &ListNode{Next: head}
    node := dummy
    for node != nil {
        next := node.Next
        for next != nil && next.Val == val {
            next = next.Next
        }
        node.Next = next
        node = next
    }
    return dummy.Next
}
