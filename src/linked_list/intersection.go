func GetIntersectionNode(headA, headB *ListNode) *ListNode {
    ptrA := headA
    ptrB := headB
    for ptrA != ptrB {
        if ptrA == nil {
            ptrA = headB
        } else {
            ptrA = ptrA.Next
        }
        if ptrB == nil {
            ptrB = headA
        } else {
            ptrB = ptrB.Next
        }
    }
    return ptrA
}
