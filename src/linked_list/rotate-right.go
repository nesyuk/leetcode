func RotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    var prev *ListNode
    cur := head
    l := 0
    for cur != nil {
        prev = cur
        cur = cur.Next
        l++
    }
    k = k % l
    prev.Next = head
    for i := 0; i < l - k; i++ {
        prev = prev.Next
    }
    head = prev.Next
    prev.Next = nil
    return head
}
