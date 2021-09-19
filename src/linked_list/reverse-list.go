package linked_list

func ReverseLinkedListRecursive(head *ListNode) *ListNode {
	return swap(head, nil)
}

func swap(node *ListNode, prev *ListNode) *ListNode {
    if node == nil {
        return prev
    }
    next := node.Next
    node.Next = prev
    return swap(next, node)
}

func ReverseLinkedListIterative(head *ListNode) *ListNode {
    var prev *ListNode
    cur := head
    for cur != nil {
        next := cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }
    return prev
}
