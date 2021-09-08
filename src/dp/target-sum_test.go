package dp

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	ways := FindTargetSumWays([]int{0,0,0,0,0,0,0,0,1}, 1)
	if ways != 256 {
		t.Fatalf("expected: %v, got: %v", 256, ways)
	}
}
