package binary_search

import "testing"

func TestFindFirstBrokenBuild(t *testing.T) {
	firstBroken := FindFirstBrokenBuild([]int{0, 0, 0, 1, 1, 1})
	if firstBroken != 3 {
		t.Fatalf("expected: %v, got: %v", 3, firstBroken)
	}
	firstBroken = FindFirstBrokenBuild([]int{0, 0, 0, 0, 0, 0})
	if firstBroken != -1 {
		t.Fatalf("expected: %v, got: %v", -1, firstBroken)
	}
	firstBroken = FindFirstBrokenBuild([]int{1, 1, 1, 1, 1, 1, 1})
	if firstBroken != 0  {
		t.Fatalf("expected: %v, got: %v", 0, firstBroken)
	}

}
