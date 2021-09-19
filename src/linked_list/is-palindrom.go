package linked_list

func IsPalindrome(head *ListNode) bool {
    return verify(head, head)
}

func verify(tail *ListNode, head *ListNode) bool {
    if tail == nil {
        return true
    }
    innerCompareResult := verify(tail.Next, head)
    if !innerCompareResult {
        return false
    }
    cmp := head.Val == tail.Val
    if head.Next != nil {
        *head = *head.Next
    }
    return cmp
}
