package array

import "testing"

func TestMaxSubArrayKadane(t *testing.T) {
	maxSum := MaxSubArrayKadane([]int{5,4,-1,7,8})
	if maxSum != 23 {
		t.Fatalf("expected: %v, got: %v", 23, maxSum)
	}
        
        maxSum = MaxSubArrayKadane([]int{-2, -1})
	if maxSum != -1 {
		t.Fatalf("expected: %v, got: %v", -1, maxSum)
	}
        
        maxSum = MaxSubArrayKadane([]int{2, 1})
	if maxSum != 3 {
		t.Fatalf("expected: %v, got: %v", 3, maxSum)
	}

        maxSum = MaxSubArrayKadane([]int{1})
	if maxSum != 1 {
		t.Fatalf("expected: %v, got: %v", 1, maxSum)
	}

        maxSum = MaxSubArrayKadane([]int{-2,1,-3,4,-1,2,1,-5,4})
	if maxSum != 6 {
		t.Fatalf("expected: %v, got: %v", 6, maxSum)
	}

}
