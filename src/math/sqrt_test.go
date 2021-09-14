package sqrt

import (
        "testing"
)

var tests = []struct{data int
                     want int}{{0, 0},{1, 1}, {2, 1}, {3, 1}, {4, 2}, {10, 3}, {2147483647, 46340}}

func TestIntSqrtBinarySearch(t *testing.T) {
	for _, test := range tests {
		sqrt := IntSqrtBinarySearch(test.data)
		if sqrt != test.want {
			t.Fatalf("expected: %v, got: %v", test.want, sqrt)
		}
	}
}

func TestSqrtPocketCalculator(t *testing.T) {
	for _, test := range tests {
		sqrt := SqrtPocketCalculator(test.data)
		if sqrt != test.want {
			t.Fatalf("expected: %v, got: %v", test.want, sqrt)
		}
	}
}

func TestSqrtRecursiveBitShifts(t *testing.T) {
	for _, test := range tests {
		sqrt := SqrtRecursiveBitShifts(test.data)
		if sqrt != test.want {
			t.Fatalf("expected: %v, got: %v", test.want, sqrt)
		}
	}
}

func TestSqrtNewtonsMethod(t *testing.T) {
	for _, test := range tests {
		sqrt := SqrtNewtonsMethod(test.data)
		if sqrt != test.want {
			t.Fatalf("expected: %v, got: %v", test.want, sqrt)
		}
	}

}
