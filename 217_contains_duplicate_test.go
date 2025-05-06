package leetcode

/*
	Given an integer array nums,
	return true if any value appears at least twice in the array, and return false if every element is distinct.

	Example 1:
	Input: nums = [1,2,3,1]
	Output: true

	Explanation:
	The element 1 occurs at the indices 0 and 3.

	Example 2:
	Input: nums = [1,2,3,4]
	Output: false

	Explanation:
	All elements are distinct.

	Example 3:

	Input: nums = [1,1,1,3,3,4,3,2,4,2]
	Output: true

	Constraints:

	1 <= nums.length <= 10^5
	-10^9 <= nums[i] <= 10^9
*/

import (
	"slices"
	"testing"
)

func containsDuplicate(nums []int) bool {
	slices.Sort(nums)
	for i := range len(nums) - 1 {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

func TestContainesDuplicate(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{
			input:    []int{1, 2, 3, 1},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 4},
			expected: false,
		},
		{
			input:    []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
	}

	for _, test := range tests {
		res := containsDuplicate(test.input)
		if res != test.expected {
			t.Errorf("res = %v, expected: %v\n", res, test.expected)
		}
	}
}

/*
	First solution
	Add each element to the map. If the key already exists, the array contains a duplicate.
	But this solution works slower.
	It is faster to sort the array first.
	If two consecutive elements are equal, then the array contains a duplicate.

	func containsDuplicate(nums []int) bool {
		m := make(map[int]struct{}, len(nums))
		for _, val := range nums {
			if _, ok := m[val]; !ok {
			m[val] = struct{}{}
				continue
			}

			return true
		}

		return false
	}
*/
