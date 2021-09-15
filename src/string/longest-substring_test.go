package string

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct{data string; want int}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
	}
	for _, test := range tests {
		subStrLen := LengthOfLongestSubstring(test.data)
		if subStrLen != test.want {
			t.Fatalf("expected: %d, got: %d", test.want, subStrLen)
		}
	}
}
