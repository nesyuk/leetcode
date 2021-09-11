package array

import "testing"

func TestMaxSubArrayDivideConquer(t *testing.T) {
	maxSum := MaxSubArrayDivideConquer([]int{5,4,-1,7,8})
	if maxSum != 23 {
		t.Fatalf("expected: %v, got: %v", 23, maxSum)
	}
        
        maxSum = MaxSubArrayDivideConquer([]int{-2, -1})
	if maxSum != -1 {
		t.Fatalf("expected: %v, got: %v", -1, maxSum)
	}
        
        maxSum = MaxSubArrayDivideConquer([]int{2, 1})
	if maxSum != 3 {
		t.Fatalf("expected: %v, got: %v", 3, maxSum)
	}

        maxSum = MaxSubArrayDivideConquer([]int{1})
	if maxSum != 1 {
		t.Fatalf("expected: %v, got: %v", 1, maxSum)
	}

        maxSum = MaxSubArrayDivideConquer([]int{-2,1,-3,4,-1,2,1,-5,4})
	if maxSum != 6 {
		t.Fatalf("expected: %v, got: %v", 6, maxSum)
	}

}
