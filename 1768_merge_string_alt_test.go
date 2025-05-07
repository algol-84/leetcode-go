package leetcode

/*
	You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1.
	If a string is longer than the other, append the additional letters onto the end of the merged string.

	Return the merged string.

	Example 1:

	Input: word1 = "abc", word2 = "pqr"
	Output: "apbqcr"
	Explanation: The merged string will be merged as so:
	word1:  a   b   c
	word2:    p   q   r
	merged: a p b q c r

	Example 2:

	Input: word1 = "ab", word2 = "pqrs"
	Output: "apbqrs"
	Explanation: Notice that as word2 is longer, "rs" is appended to the end.
	word1:  a   b
	word2:    p   q   r   s
	merged: a p b q   r   s

	Example 3:

	Input: word1 = "abcd", word2 = "pq"
	Output: "apbqcd"
	Explanation: Notice that as word1 is longer, "cd" is appended to the end.
	word1:  a   b   c   d
	word2:    p   q
	merged: a p b q c   d

	Constraints:
	1 <= word1.length, word2.length <= 100
	word1 and word2 consist of lowercase English letters.
*/

import (
	"strings"
	"testing"
)

func mergeAlternately(word1 string, word2 string) string {
	var res strings.Builder
	res.Grow(len(word1) + len(word2))

	p1, p2 := 0, 0
	for p1 < len(word1) && p2 < len(word2) {
		res.WriteByte(word1[p1])
		res.WriteByte(word2[p2])
		p1++
		p2++
	}

	res.WriteString(word1[p1:])
	res.WriteString(word2[p2:])

	return res.String()
}

func Test_mergeAlternately(t *testing.T) {
	tests := []struct {
		word1    string
		word2    string
		expected string
	}{
		{
			"abc",
			"pqr",
			"apbqcr",
		},
		{
			"ab",
			"pqrs",
			"apbqrs",
		},
		{
			"abcd",
			"pq",
			"apbqcd",
		},
	}

	for _, test := range tests {
		res := mergeAlternately(test.word1, test.word2)
		if res != test.expected {
			t.Errorf("merged string = %s; want %s", res, test.expected)
			continue
		}

		t.Logf("merged string %s :ok", res)
	}

}

/*
	The solution using slice is slower than using builder

	func mergeAlternately(word1 string, word2 string) string {
		res := make([]byte, len(word1)+len(word2))

		var length byte
		if len(word1) < len(word2) {
			length = byte(len(word1))
			copy(res[len(word1)*2:], word2[len(word1):])
		} else {
			length = byte(len(word2))
			copy(res[len(word2)*2:], word1[len(word2):])
		}

		for i := range length {
			res[2*i] = word1[i]
			res[2*i+1] = word2[i]
		}

		return string(res)
	}

*/
