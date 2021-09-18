package linked_list

import "testing"

func TestFloydCyclePositionDetection(t *testing.T) {
	tests := []struct{data *ListNode; want *ListNode}{
		{NewLinkedList([]int{}), nil},
		{NewLinkedList([]int{1}), nil},
		{NewLinkedListWithCycle([]int{1}, 0), &ListNode{Val: 1}},
		{NewLinkedListWithCycle([]int{1,2}, 0), &ListNode{Val: 1}},
		{NewLinkedListWithCycle([]int{3, 2, 0, -4}, 0), &ListNode{Val: 3}},
		{NewLinkedListWithCycle([]int{3, 2, 0, -4}, 1), &ListNode{Val: 2}},
		{NewLinkedListWithCycle([]int{3, 2, 0, -4}, 2), &ListNode{Val: 0}},
		{NewLinkedListWithCycle([]int{3, 2, 0, -4}, 3), &ListNode{Val: -4}},
		{NewLinkedListWithCycle([]int{1,2,3,4}, 1), &ListNode{Val: 2}},
	}
	for _, test := range tests {
		node := GetCyclePosition(test.data)
		if test.want == nil && node != nil {
			t.Fatalf("expected: %v, got: %v", test.want, node.Val)
		} else if test.want != nil && node.Val != test.want.Val {
			t.Fatalf("expected: %v, got: %v", test.want.Val, node.Val)
		}
	}
}
