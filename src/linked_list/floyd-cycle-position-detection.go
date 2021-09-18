package linked_list


func GetCyclePosition(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	intersection := GetIntersection(head)
	if intersection == nil {
		return nil
	}
	ptr1 := head
	ptr2 := intersection

	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	return ptr1
}

func GetIntersection(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}
	return nil
}
