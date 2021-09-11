package bstree

import (
	"testing"
)

func TestNewTree(t *testing.T) {
	root := NewBSTree([]int{5,3,6,2,4,1})
	verify(t, InorderSuccessor(root, root.Find(5)).Val, 6)
	verify(t, InorderSuccessor(root, root.Find(3)).Val, 4)
	successor := InorderSuccessor(root, root.Find(6))
	if successor != nil {
            t.Fatalf("expected: %v, got: %v", nil, successor)
	}
	verify(t, InorderSuccessor(root, root.Find(2)).Val, 3)
	verify(t, InorderSuccessor(root, root.Find(4)).Val, 5)
	verify(t, InorderSuccessor(root, root.Find(1)).Val, 2)
}

func verify(t *testing.T, got int, expected int) {
	if expected != got {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}
}
