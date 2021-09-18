package linked_list

import "testing"

func TestFloydCycleDetection(t *testing.T) {
	tests := []struct{data *ListNode; want bool}{
		{NewLinkedListWithCycle([]int{1,2}, 0), true},
		{NewLinkedListWithCycle([]int{3,2,0,-4}, 2), true},
		{NewLinkedListWithCycle([]int{3,2,0,-4}, 1), true},
		{NewLinkedListWithCycle([]int{3,2,0,3,-4}, 1), true},
		{NewLinkedListWithCycle([]int{3,2,0,3,4,5,7-4}, 1), true},
		{NewLinkedListWithCycle([]int{3,2,-4}, 1), true},
		{NewLinkedList([]int{3,2,0,-4}), false},
	}
	for _, test := range tests {
		ans := HasCycle(test.data)
		if ans != test.want {
			t.Fatalf("expected: %v, got: %v", test.want, ans)
		}
	}
}
