package bit_manipulation

import "testing"

func TestKthGrammar(t *testing.T) {
	tests := []struct{
		data []int
		want int}{
		{[]int{1, 1}, 0},
		{[]int{8, 2}, 1},
	}
	for _, test := range tests {
		bit := KthGrammar(test.data[0], test.data[1])
		if bit !=  test.want {
			t.Fatalf("expected: %v, got: %v", test.want, bit)
		}
	}
}
