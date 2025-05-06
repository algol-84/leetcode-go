/*
Write a function that reverses a string. The input string is given as an array of characters s.
You must do this by modifying the input array in-place with O(1) extra memory.

Example 1:

Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
Example 2:

Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]

Constraints:

1 <= s.length <= 10^5
s[i] is a printable ascii character.
*/
package leetcode

import (
	"testing"
)

func reverseString(s []byte) {
	for i := range len(s) / 2 {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"hello",
			"olleh",
		},
		{
			"Hannah",
			"hannaH",
		},
		{
			"Go",
			"oG",
		},
	}

	for _, test := range tests {
		bytes := []byte(test.input)
		reverseString(bytes)
		if string(bytes) != test.expected {
			t.Errorf("reverseString(%q) = %q; want %q", test.input, string(bytes), test.expected)
			continue
		}

		t.Logf("%q --> %q :ok", test.input, string(bytes))
	}

}
