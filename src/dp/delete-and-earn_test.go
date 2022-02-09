package dp

import "testing"


func TestDeleteAndEarn(t *testing.T) {
    input := []int{8,3,4,7,6,6,9,2,5,8,2,4,9,5,9,1,5,7,1,4}

    if res := DeleteAndEarn(input); res != 61 {
        t.Fatalf("expected: %d, got: %d", 61, res)
    }
}
