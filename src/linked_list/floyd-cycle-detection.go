package linked_list


func HasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }

    slow := head
    fast := head.Next

    for slow != fast {
        if (fast == nil || fast.Next == nil) {
            return false
        }
        slow, fast = slow.Next, fast.Next.Next
    }
    return true
}
