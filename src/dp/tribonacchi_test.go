package dp

import "testing"


func Test_TribonacciTopDown(t *testing.T) {
    tests := []struct{n int; expect int} {
       {0, 0},
       {1, 1},
       {2, 1},
       {3, 2},
       {4, 4},
       {37, 2082876103},
    }
    for _, test := range tests {
        if res := TribonacciTopDown(test.n); res != test.expect {
            t.Fatalf("expected: %d, got: %d", test.expect, res);
        }
    }
}

func Test_TribonacciBottomUp(t *testing.T) {
    tests := []struct{n int; expect int} {
       {0, 0},
       {1, 1},
       {2, 1},
       {3, 2},
       {4, 4},
       {37, 2082876103},
    }
    for _, test := range tests {
        if res := TribonacciBottomUp(test.n); res != test.expect {
            t.Fatalf("expected: %d, got: %d", test.expect, res);
        }
    }
}
